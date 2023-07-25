package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List_opening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
