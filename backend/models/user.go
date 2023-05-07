package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User is a struct that represents a user in the database
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username    string             `bson:"username" json:"username" unique:"true"`
	DisplayName string             `bson:"display_name,omitempty" json:"display_name,omitempty"`
	Salt        string             `bson:"salt" json:"salt"`
	Password    string             `bson:"password" json:"password"`
}
