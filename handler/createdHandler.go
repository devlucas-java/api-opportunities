package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "200 Created",
	})
}
