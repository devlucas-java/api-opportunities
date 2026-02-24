package router

import "github.com/gin-gonic/gin"

func Initalizer() {

	// Initalizer router
	router := gin.Default()
	// Initalizer routes
	initalizerRoutes(router)
	// Run the server
	router.Run(":8080")
}
