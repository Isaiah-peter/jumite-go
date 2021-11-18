package models

import (
	"awesomeProject/jumite/pkg/config"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ConversationId int64  `json:"conversation_id"`
	Text           string `json:"text"`
	SenderId       int64  `json:"sender_id"`
	Conversation   Conversation
}

type Conversation struct {
	gorm.Model
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	RecieverId int64 `json:"reciever_id"`
	Message    []Message
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Message{})
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Message{})
}
