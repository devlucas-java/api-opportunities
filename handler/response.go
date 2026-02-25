package handler

import (
	"fmt"
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(cxt *gin.Context, status int, message string) {
	cxt.Header("Content-Type", "application/json")
	cxt.JSON(status, gin.H{
		"message": message,
		"status":  status,
	})
}

func SendSuccessResponse(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success in the operation: %s", operation),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
