package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID 		int
	Name 	string
	Check	int
}
