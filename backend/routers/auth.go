package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/controllers"
)

func AuthRouter(router *gin.Engine, userController *controllers.UserController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/register", userController.Register)
	}
}
