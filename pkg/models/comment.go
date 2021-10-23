package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	Text         string `json:"text"`
}
