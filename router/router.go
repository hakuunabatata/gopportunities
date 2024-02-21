package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	initRoutes(router)

	port := ":3000"

	router.Run(port)
}
