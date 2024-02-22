package router

import (
	gin "github.com/gin-gonic/gin"
	handler "github.com/hakuunabatata/gopportunities/handler"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opportunity", handler.ShowOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpportunityHandler)
		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)
		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
		v1.GET("/opportunities", handler.ListOpportunityHandler)
	}
}
