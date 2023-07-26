package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelc484/go-estudo.git/schemas"
)

func List_opening_handler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		send_error(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	send_success(ctx, "list-openings", openings)
}
