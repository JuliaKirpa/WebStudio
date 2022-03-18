package main

import (
	"ClientBaseWEBStudio/api"
	"ClientBaseWEBStudio/pkg/handler"
	"log"
)

func main() {
	handler := new(handler.Handler)

	server := new(api.Server)
	err := server.Run("8080", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error to running server %s", err)
	}

}
