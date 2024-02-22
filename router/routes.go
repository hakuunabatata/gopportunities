package router

import (
	gin "github.com/gin-gonic/gin"
	handler "github.com/hakuunabatata/gopportunities/handler"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opportunity", handler.ShowOpeningHandler)
		v1.POST("/opportunity", handler.CreateOpeningHandler)
		v1.PUT("/opportunity", handler.UpdateOpeningHandler)
		v1.DELETE("/opportunity", handler.DeleteOpeningHandler)
		v1.GET("/opportunities", handler.ListOpeningHandler)
	}
}
