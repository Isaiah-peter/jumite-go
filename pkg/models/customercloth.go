package models

import (
	"gorm.io/gorm"
)

type CustomerProduct struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	Image        string `json:"img"`
}

func init() {
	db.AutoMigrate(&CustomerProduct{})
}
