package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuchia0221/Chatify/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *UserController) Login(ctx *gin.Context) {
	var body struct {
		Username string
		Password string
	}

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

	filter := bson.M{"username": body.Username}
	var user models.User

	err := c.Collection.FindOne(context.Background(), filter).Decode(&user)
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

	if !VerifyPassword(body.Password, user.Salt, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Create the JWT token
	token, err := GenerateToken(body.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	// Set the JWT token as a cookie for the client
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
	})
}

func (c *UserController) Register(ctx *gin.Context) {
	var body struct {
		Username    string
		DisplayName string
		Password    string
	}
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

	if body.DisplayName == "" {
		body.DisplayName = body.Username
	}

	var user models.User
	filter := bson.M{"username": body.Username}
	err := c.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}

	user.Username, user.DisplayName = body.Username, body.DisplayName
	user.Salt, user.Password, err = HashPassword(body.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	result, err := c.Collection.InsertOne(context.Background(), user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":      result.InsertedID,
		"message": "User created successfully",
	})
}
