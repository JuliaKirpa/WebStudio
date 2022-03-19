package main

import (
	"ClientBaseWEBStudio/api"
	"ClientBaseWEBStudio/pkg/handler"
	"ClientBaseWEBStudio/pkg/repository"
	"ClientBaseWEBStudio/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error from init config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(api.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
