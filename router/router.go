package router

import "github.com/gin-gonic/gin"

func Initializer() {

	// Initalizer router
	router := gin.Default()
	// Initalizer routes
	initializerRoutes(router)
	// Run the server
	router.Run(":8080")
}
