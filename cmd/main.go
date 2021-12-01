package main

import (
	"ClientBaseWEBStudio/api"
	"ClientBaseWEBStudio/pkg/repository"
)

func main(){
	api.ServerUp()
	repository.DBconnect()

}
