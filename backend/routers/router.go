package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/controllers"
	"github.com/yuchia0221/Chatify/database"
)

// InitRouter initializes the gin router and registers the routes
func InitRouter() *gin.Engine {
	router := gin.Default()

	// Define the routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	userCollection := database.GetMongoClient().Database("db").Collection("users")
	userController := &controllers.UserController{Collection: userCollection}
	AuthRouter(router, userController)

	return router
}
