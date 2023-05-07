package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitMongoDB initializes the MongoDB client and returns a reference to the client
func InitMongoDB() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
}

func GetMongoClient() *mongo.Client {
	return client
}
