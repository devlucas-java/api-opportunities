package router

import (
	"github.com/devlucas-java/api-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func initalizerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowHandler)
		v1.PUT("/opening", handler.UpdateHandler)
		v1.POST("/opening", handler.CreatedHandler)
		v1.DELETE("/opening", handler.DeleteHandler)
		v1.GET("/openings", handler.ListHandler)
	}
}
