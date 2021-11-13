package models

import (
	"awesomeProject/jumite/pkg/config"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	OrderStatus string `json:"order_status" asn1:"default:pending"`
	Delete      string `json:"delete"`
	Messages    []Message
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Order{})
}
