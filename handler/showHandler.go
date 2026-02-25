package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Show an opening
// @Description Show an opening by its ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse "Successfully retrieved opening"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "Opening not found"
// @Router /opening [get]
func ShowHandler(ctx *gin.Context) {

	opening := schemas.Opening{}

	id := ctx.Query("id")
	if id == "" {
		SendErrorResponse(
			ctx,
			http.StatusBadRequest,
			errParamRequerided("id", "string").Error())
		return
	}
	if err := db.Find(&opening, id).Error; err != nil {
		SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	SendSuccessResponse(ctx, "show-opening", opening)
}
