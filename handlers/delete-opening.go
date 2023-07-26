package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelc484/go-estudo.git/schemas"
)

func Delete_opening_handler(ctx *gin.Context) {
	id := ctx.Query("id")

	fmt.Println(id)

	if id == "" {
		send_error(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	err := db.Where("id = ?", id).First(&opening).Error

	fmt.Println(err)

	// Find opening
	if err != nil {
		send_error(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		send_error(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	send_success(ctx, "delete-opening", opening)
}
