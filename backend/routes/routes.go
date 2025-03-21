package routes

import (
	"github.com/gabrielg2020/portfolio/backend/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API routes
	api := router.Group("/api")
	{
		api.GET("/hello", handlers.HelloHandler)
	}

	// Serve static files
	router.Static("/assets", "./backend/static/assets")

	// Serve index.html for any unmatched routes
	router.NoRoute(func(ctx *gin.Context) {
		// Check if the request is for an API route
		if path.Ext(ctx.Request.URL.Path) == "" && ctx.Request.URL.Path[:4] != "/api" {
			ctx.File("./backend/static/index.html")
		} else {
			ctx.Status(http.StatusNotFound)
		}
	})

	return router
}
