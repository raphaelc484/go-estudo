package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelc484/go-estudo.git/schemas"
)

func Update_opening_handler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		send_error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		send_error(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schemas.Opening{}

	err := db.Where("id = ?", id).First(&opening).Error

	if err != nil {
		send_error(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("error updating opening: %v", err.Error())
		send_error(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	send_success(ctx, "update-opening", opening)
}
