package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelc484/go-estudo.git/handlers"
)

func initialize_routes(router *gin.Engine) {
	handlers.Initialize_handler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.Show_opening_handler)
		v1.POST("/opening", handlers.Create_opening_handler)
		v1.DELETE("/opening", handlers.Delete_opening_handler)
		v1.PUT("/opening", handlers.Update_opening_handler)
		v1.GET("/openings", handlers.List_opening_handler)
	}
}
