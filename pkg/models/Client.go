package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID 		int
	Name 	string
	Check	int
}

