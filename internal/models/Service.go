package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ID    uint   `gorm:"primary_key; auto-increment" json:"ID"`
	Name  string `gorm:"size 255; not-null" json:"name"`
	Price uint   `json:"price"`
}
