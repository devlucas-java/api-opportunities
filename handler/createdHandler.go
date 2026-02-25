package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreatedHandler(ctx *gin.Context) {

	request := CreatedOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error in validate request: %v", err.Error())
		SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Remote:   *request.Remote,
		Salary:   request.Salary,
		Location: request.Location,
		Link:     request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error in create record: %v", err.Error())
		SendErrorResponse(ctx, http.StatusInternalServerError, "Error in create record")
		return
	}

	SendSuccessResponse(ctx, "created-opening", opening)
}
