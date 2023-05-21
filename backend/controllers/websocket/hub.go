package websocket

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yuchia0221/Chatify/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HubController struct {
	Hub     *models.Hub
	Room    *mongo.Collection
	Message *mongo.Collection
}

type CreateRoomReq struct {
	RoomName string `json:"name"`
}

type RoomRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ClientNum int    `json:"client_num"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewHubController(h *models.Hub, RoomCollection *mongo.Collection, MessageCollection *mongo.Collection) *HubController {
	cursor, err := RoomCollection.Find(context.Background(), bson.M{})
	if err == nil {
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var room models.RoomData
			cursor.Decode(&room)
			h.Rooms[room.ID.Hex()] = &models.Room{
				ID:      room.ID.Hex(),
				Name:    room.Name,
				Clients: make(map[string]*models.Client),
			}
		}
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}

	return &HubController{
		Hub:     h,
		Room:    RoomCollection,
		Message: MessageCollection,
	}
}

func (h *HubController) CreateRoom(ctx *gin.Context) {
	var body CreateRoomReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.RoomName == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	var room models.RoomData
	filter := bson.M{"name": body.RoomName}
	err := h.Room.FindOne(context.Background(), filter).Decode(&room)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "Room name already exists",
		})
		return
	}

	room.Name, room.Clients = body.RoomName, []string{}
	result, err := h.Room.InsertOne(context.Background(), room)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	h.Hub.Rooms[id] = &models.Room{
		ID:      id,
		Name:    body.RoomName,
		Clients: make(map[string]*models.Client),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"room_id": id,
		"message": "Room created successfully",
	})
}

func (h *HubController) JoinRoom(ctx *gin.Context) {
	roomID := ctx.Param("roomId")
	username := ctx.MustGet("username").(string)
	if roomID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	// if client already exists in room, return error
	if _, ok := h.Hub.Rooms[roomID].Clients[username]; ok {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "Client already exists in room",
		})
		return
	}

	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid room id",
		})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$addToSet": bson.M{"clients": username}}
	opts := options.Update().SetUpsert(true)
	_, err = h.Room.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add username to room",
		})
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upgrade connection to WebSocket",
		})
		return
	}

	client := &models.Client{
		Conn:     conn,
		Message:  make(chan *models.Message, 10),
		RoomID:   roomID,
		Username: username,
	}

	message := &models.Message{
		Content:  username + " joined the room",
		RoomID:   roomID,
		Username: username,
		SendTime: time.Now(),
	}

	h.Hub.Register <- client
	h.Hub.Broadcast <- message

	clientController := NewClientController(client, h.Message)

	go clientController.wsWriteMessage()
	clientController.wsReadMessage(h.Hub)
}

func (h *HubController) LeaveRoom(ctx *gin.Context) {
	roomID := ctx.Param("roomId")
	username := ctx.MustGet("username").(string)
	if roomID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	if _, ok := h.Hub.Rooms[roomID]; !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Room not found",
		})
		return
	}

	if _, ok := h.Hub.Rooms[roomID].Clients[username]; !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Client not found in room",
		})
		return
	}

	client := h.Hub.Rooms[roomID].Clients[username]
	h.Hub.Unregister <- client

	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid room id",
		})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$pull": bson.M{"clients": username}}
	result, err := h.Room.UpdateOne(context.Background(), filter, update)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove username from room",
		})
		return
	}

	if result.ModifiedCount == 0 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove username from room",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Left room successfully",
	})
}

func (h *HubController) GetAllRooms(ctx *gin.Context) {
	var rooms []RoomRes
	cursor, err := h.Room.Find(context.Background(), bson.M{})
	if err == nil {
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var room models.RoomData
			cursor.Decode(&room)
			rooms = append(rooms, RoomRes{ID: room.ID.Hex(), Name: room.Name, ClientNum: len(room.Clients)})
		}
	}

	if err := cursor.Err(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to query database",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}

// get all clients in a room
func (h *HubController) GetAllClientsInRoom(ctx *gin.Context) {
	roomID := ctx.Param("roomId")
	if roomID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid room id",
		})
		return
	}

	var room models.RoomData
	filter := bson.M{"_id": id}
	err = h.Room.FindOne(context.Background(), filter).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Room not found",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to query database",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"clients": room.Clients,
	})
}
