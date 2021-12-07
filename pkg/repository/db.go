package repository

import (
	"ClientBaseWEBStudio/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DBconnect(){
	dsn := "host=localhost user=gipernova password=qwerty dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Client{})


	db.Create(&models.Client{ID: 6778, Name: "Lia", Check: 1456})

	fmt.Println("Connect to pg success")
}


