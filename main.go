package main

import (
	"github.com/raphaelc484/go-estudo.git/config"
	"github.com/raphaelc484/go-estudo.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()

}
