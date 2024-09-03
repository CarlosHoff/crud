package main

import (
	"fmt"

	"github.com/CarlosHoff/crud.git/config"
	"github.com/CarlosHoff/crud.git/router"
)

var logger *config.Logger

func main() {

	//Init config
	err := config.Init()

	if err != nil {
		fmt.Println("ERRO DB")
		logger.Errf("config initialize error: %v", err)
		return
	}

	//Init Routers
	router.Initialize()
}
