package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID        uint   `gorm:"primary_key; auto-increment" json:"id"`
	Name      string `gorm:"size 255; not-null" json:"name"`
	Telephone int    `gorm:"varchar(50); not-null" json:"number"`
	Email     string `gorm:"size 255" json:"email"`
}
