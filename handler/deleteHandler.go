package handler

import (
	"fmt"
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		SendErrorResponse(
			ctx,
			http.StatusBadRequest,
			errParamRequerided("id", "string").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("Opening it id: %v not found", id))
		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Error delete opening id: %v", id))
		return
	}
	SendSuccessResponse(ctx, fmt.Sprintf("delete opening id: %v"+id), opening)
}
