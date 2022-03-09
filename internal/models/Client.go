package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	ID        uint      `gorm:"primary_key; auto-increment" json:"id"`
	Name      string    `gorm:"size 255; not-null" json:"name"`
	Telephone int       `gorm:"varchar(50); not-null" json:"number"`
	Email     string    `gorm:"size 255" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
