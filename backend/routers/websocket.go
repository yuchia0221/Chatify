package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/controllers/websocket"
	"github.com/yuchia0221/Chatify/middlewares"
)

func WebSocketRouter(router *gin.Engine, hubController *websocket.HubController) {
	websocketGroup := router.Group("/websocket")
	websocketGroup.Use(middlewares.JWTAuthMiddleware())
	{
		websocketGroup.POST("/room", hubController.CreateRoom)
		websocketGroup.GET("/room/:roomId", hubController.GetAllClientsInRoom)
		websocketGroup.DELETE("/room/:roomId", hubController.LeaveRoom)
		websocketGroup.GET("/joinRoom/:roomId", hubController.JoinRoom)
		websocketGroup.GET("/rooms", hubController.GetAllRooms)
	}
}
