package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary List all openings
// @Description Retrieve a list of all openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse "Successfully retrieved list of openings"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /openings [get]
func ListHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve list of openings")
		return
	}
	SendSuccessResponse(ctx, "list-openings", openings)
}
