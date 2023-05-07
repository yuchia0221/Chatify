package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/controllers"
	"github.com/yuchia0221/Chatify/middlewares"
)

func UserRouter(router *gin.Engine, userController *controllers.UserController) {
	userGroup := router.Group("/user")
	userGroup.Use(middlewares.JWTAuthMiddleware())
	{
		userGroup.GET("/", userController.GetUser)
		userGroup.PUT("/display_name", userController.UpdateDisplayName)
		userGroup.PUT("/password", userController.UpdatePassword)
		userGroup.DELETE("/", userController.DeleteUser)
	}
}
