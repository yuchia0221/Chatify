package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserController is a struct that holds the necessary information
// to interact with the user model
type UserController struct {
	User *mongo.Collection
}

type UpdateDisplayNameReq struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type UpdatePasswordReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserController(UserCollection *mongo.Collection) *UserController {
	return &UserController{
		User: UserCollection,
	}
}

// GetUser is a function that gets a user's display name from the database
func (c *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		username = ctx.MustGet("username").(string)
	}

	filter := bson.M{"username": username}

	var user models.User

	err := c.User.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to query database",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"username":     user.Username,
		"display_name": user.DisplayName,
	})
}

// UpdateDisplayName is a function that updates a user's display name in the database
func (c *UserController) UpdateDisplayName(ctx *gin.Context) {
	var body UpdateDisplayNameReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.Username == "" || body.DisplayName == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	// get document from database
	username := ctx.MustGet("username").(string)
	if username != body.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"display_name": body.DisplayName}}

	result, err := c.User.UpdateOne(context.Background(), filter, update)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user document",
		})
		return
	}
	if result.MatchedCount == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Display name updated successfully",
	})
}

// UpdatePassword is a function that updates a user's password in the database
func (c *UserController) UpdatePassword(ctx *gin.Context) {
	var body UpdatePasswordReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.Username == "" || body.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	username := ctx.MustGet("username").(string)
	if username != body.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	salt, password, err := HashPassword(body.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"salt": salt, "password": password}}

	result, err := c.User.UpdateOne(context.Background(), filter, update)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user document",
		})
		return
	}

	if result.MatchedCount == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}

// DeleteUser is a function that deletes a user from the database
func (c *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)

	filter := bson.M{"username": username}
	result, err := c.User.DeleteOne(context.Background(), filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user document",
		})
		return
	}
	if result.DeletedCount == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
