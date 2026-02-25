package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	opening := schemas.Opening{}
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	// validate query id
	if id == "" {
		SendErrorResponse(
			ctx,
			http.StatusNotFound,
			errParamRequerided("id", "string").Error())
		return
	}
	// validate request
	if err := request.Validate(); err != nil {
		logger.Errorf("Error in validate request: %v", err.Error())
		SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	// find record by id
	if err := db.Find(&opening, id).Error; err != nil {
		SendErrorResponse(ctx, http.StatusFound, err.Error())
		return
	}
	// update record
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// save record
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error in save record: %v", err.Error())
		SendErrorResponse(ctx, http.StatusInternalServerError, "Error in save record")
		return
	}
	SendSuccessResponse(ctx, "save-record", opening)
}
