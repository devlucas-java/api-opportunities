package router

import (
	"github.com/devlucas-java/api-opportunities/docs"
	"github.com/devlucas-java/api-opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializerRoutes(router *gin.Engine) {

	// Initialize the handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowHandler)
		v1.PUT("/opening", handler.UpdateHandler)
		v1.POST("/opening", handler.CreatedHandler)
		v1.DELETE("/opening", handler.DeleteHandler)
		v1.GET("/openings", handler.ListHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
