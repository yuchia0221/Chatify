package main

import (
	"context"

	"github.com/yuchia0221/Chatify/config"
	"github.com/yuchia0221/Chatify/database"
	"github.com/yuchia0221/Chatify/routers"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize MongoDB client
	database.InitMongoDB()
	client := database.GetMongoClient()
	defer client.Disconnect(context.Background())

	// Initialize the router
	router := routers.InitRouter()
	router.Run()
}
