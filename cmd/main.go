package main

import (
	"ClientBaseWEBStudio/api"
	"ClientBaseWEBStudio/pkg/handler"
	"ClientBaseWEBStudio/pkg/repository"
	"ClientBaseWEBStudio/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(api.Server)
	err := server.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error to running server %s", err)
	}

}
