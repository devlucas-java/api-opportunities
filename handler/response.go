package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(cxt *gin.Context, statusCode int, message string) {
	cxt.Header("Content-Type", "application/json")
	cxt.JSON(statusCode, gin.H{
		"message":     message,
		"status code": statusCode,
	})
}

func SendSuccessResponse(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success in the operation: %s", operation),
		"data":    data,
	})
}
