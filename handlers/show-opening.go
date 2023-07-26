package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelc484/go-estudo.git/schemas"
)

func Show_opening_handler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		send_error(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	err := db.Where("id = ?", id).First(&opening).Error

	if err != nil {
		send_error(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	send_success(ctx, "show-opening", opening)
}
