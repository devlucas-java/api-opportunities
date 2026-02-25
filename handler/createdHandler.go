package handler

import (
	"net/http"

	"github.com/devlucas-java/api-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new opening
// @Description Create a new opening with the provided details
// @Tags Openings
// @Accept json
// @Produce json
// @Param opening body CreatedOpeningRequest true "Opening details"
// @Success 200 {object} CreateOpeningResponse "Successfully created opening"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /opening [post]
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
