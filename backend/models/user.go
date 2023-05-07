package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User is a struct that represents a user in the database
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username    string             `json:"username"`
	DisplayName string             `json:"display_name"`
	Salt        string             `json:"salt"`
	Password    string             `json:"password"`
}
