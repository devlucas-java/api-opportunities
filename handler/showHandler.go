package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

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
