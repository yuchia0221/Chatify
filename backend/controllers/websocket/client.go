package websocket

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yuchia0221/Chatify/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientController struct {
	Client  *models.Client
	Message *mongo.Collection
}

func NewClientController(client *models.Client, MessageCollection *mongo.Collection) *ClientController {
	return &ClientController{
		Client:  client,
		Message: MessageCollection,
	}
}

func (c *ClientController) wsWriteMessage() {
	defer func() {
		c.Client.Conn.Close()
	}()

	for {
		message, ok := <-c.Client.Message
		if !ok {
			return
		}

		go c.WriteMessageToDb(message)
		c.Client.Conn.WriteJSON(message)
	}
}

func (c *ClientController) WriteMessageToDb(message *models.Message) {
	if c.Client.Username != message.Username {
		return
	}

	id, err := primitive.ObjectIDFromHex(message.RoomID)
	if err != nil {
		log.Println("Failed to convert room id to object id: ", err)
		return
	}

	messageData := &models.MessageData{
		Content:  message.Content,
		RoomID:   id,
		Username: message.Username,
		SendTime: message.SendTime,
	}

	_, err = c.Message.InsertOne(context.Background(), messageData)
	if err != nil {
		log.Println("Failed to insert message to database: ", err)
	}
}

func (c *ClientController) wsReadMessage(hub *models.Hub) {
	defer func() {
		hub.Unregister <- c.Client
		c.Client.Conn.Close()
	}()

	for {
		_, message, err := c.Client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error: ", err)
			}
			break
		}

		broadcastMessage := &models.Message{
			Content:  string(message),
			RoomID:   c.Client.RoomID,
			Username: c.Client.Username,
			SendTime: time.Now(),
		}

		hub.Broadcast <- broadcastMessage
	}
}
