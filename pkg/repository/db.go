package repository

import (
	"ClientBaseWEBStudio/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func DBconnect(){
	dsn := "host=localhost user=gipernova password=qwerty dbname=cbwebs port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = models.InitModels(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to pg success")

}


