package routes

import (
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gabrielg2020/backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API routes
	api := router.Group("/api")
	{
		api.GET("/hello", handlers.HelloHandler)
	}

	// Serve static files
	router.Static("/assets", "./static/assets")
	router.Static("/images", "./static/images")
	

	// Serve index.html for any unmatched routes
	router.NoRoute(func(ctx *gin.Context) {
		reqPath := ctx.Request.URL.Path

		// Check if the static directory exists and if index.html is there
		indexPath := "./static/index.html"
		absPath, _ := filepath.Abs(indexPath)
		log.Printf("Looking for index.html at: %s", absPath)

		// Check if the request is for an API route - safely
		isAPIPath := len(reqPath) >= 4 && reqPath[:4] == "/api"
		hasExtension := path.Ext(reqPath) != ""

		if !isAPIPath && !hasExtension {
			ctx.File(indexPath)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	})

	return router
}
