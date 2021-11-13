package models

import (
	"awesomeProject/jumite/pkg/config"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserId     int64  `json:"user_id"`
	Text       string `json:"text"`
	SenderId   int64  `json:"sender_id"`
	RecieverId int64  `json:"reciever_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Message{})
}
