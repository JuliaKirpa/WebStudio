package main

import (
	"ClientBaseWEBStudio/api"
	"ClientBaseWEBStudio/pkg/repository"
)

func main(){
	go repository.DBconnect()
	api.ServerUp()
}
