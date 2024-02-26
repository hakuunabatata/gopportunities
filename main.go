package main

import (
	fmt "fmt"

	config "github.com/hakuunabatata/gopportunities/config"
	router "github.com/hakuunabatata/gopportunities/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Init()
}
