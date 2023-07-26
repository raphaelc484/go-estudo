package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelc484/go-estudo.git/schemas"
)

func Create_opening_handler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		send_error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		send_error(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	send_success(ctx, "create-opening", opening)
}
