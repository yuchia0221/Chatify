package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/controllers"
	"github.com/yuchia0221/Chatify/controllers/websocket"
	"github.com/yuchia0221/Chatify/database"
	"github.com/yuchia0221/Chatify/middlewares"
	"github.com/yuchia0221/Chatify/models"
)

// InitRouter initializes the gin router and registers the routes
func InitRouter(hub *models.Hub) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	// Define the routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	userCollection := database.GetMongoCollection("db", "users")
	roomCollection := database.GetMongoCollection("db", "rooms")
	messageCollection := database.GetMongoCollection("db", "messages")

	userController := controllers.NewUserController(userCollection)
	UserRouter(router, userController)
	AuthRouter(router, userController)

	hubController := websocket.NewHubController(hub, roomCollection, messageCollection)
	WebSocketRouter(router, hubController)

	return router
}
