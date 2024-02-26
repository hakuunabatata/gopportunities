package main

import (
	config "github.com/hakuunabatata/gopportunities/config"
	router "github.com/hakuunabatata/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	router.Init()
}
