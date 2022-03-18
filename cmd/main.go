package main

import (
	"ClientBaseWEBStudio/api"
	"log"
)

func main() {
	server := new(api.Server)
	err := server.Run("8080");
	if err != nil {
		log.Fatalf("error to running server %s", err)
	}

}
