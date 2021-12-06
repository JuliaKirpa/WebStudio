package models

import "gorm.io/gorm"

type Client struct {
	ID 		int
	Name 	string
	Check	int
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Client{})
	if err != nil {
		return err
	}
	return nil
}
