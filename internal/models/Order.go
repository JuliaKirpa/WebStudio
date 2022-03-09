package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID        uint      `gorm:"primary_key; auto-increment" json:"id"`
	ClientID  uint      `gorm:"not null" json:"client_id"`
	Service   []Service `json:"service"`
	Done      bool      `gorm:"default:false" json:"done"`
	Discount  uint      `json:"discount"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
