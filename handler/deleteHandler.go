package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "201 No Content",
	})
}
