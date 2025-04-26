package routes

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API routes
	api := router.Group("/api")
	{
		api.GET("/version", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "HEY!",
			})
		})
	}

	// Page routes
	router.GET("/", handlers.Home)

	// Load templates
	router.LoadHTMLGlob("views/**/*.html")

	// Static files
	router.Static("/static", "./static")

	return router
}
