package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve list of openings")
		return
	}
	SendSuccessResponse(ctx, "list-openings", openings)
}
