package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username    string             `bson:"username" json:"username" unique:"true" validate:"required"`
	DisplayName string             `bson:"display_name,omitempty" json:"display_name,omitempty"`
	Salt        string             `bson:"salt" json:"salt" validate:"required"`
	Password    string             `bson:"password" json:"password" validate:"required"`
}
