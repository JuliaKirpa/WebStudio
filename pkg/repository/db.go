package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DBconnect(){
	dsn := "host=localhost user=gipernova password=qwerty dbname=cbWEBs port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = InitModels(db)
	if err != nil {
		panic(err)
	}
}

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
