package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update_opening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "UPDATE Opening",
	})
}
