package models

import (
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	RoomID   string
	Username string
}

type Message struct {
	Content  string
	RoomID   string
	Username string
	SendTime time.Time
}

type MessageData struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Content  string             `bson:"content" json:"content" validate:"required"`
	RoomID   primitive.ObjectID `bson:"room_id" json:"room_id" validate:"required"`
	Username string             `bson:"username" json:"username" validate:"required"`
	SendTime time.Time          `bson:"send_time" json:"send_time" validate:"required"`
}
