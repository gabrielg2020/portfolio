package routes

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/handlers"
	"github.com/gabrielg2020/portfolio/internal/state"
	"github.com/gin-gonic/gin"
)

func SetupRouter(state *state.State) *gin.Engine {
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
	router.GET("/", func(ctx *gin.Context) {
		handlers.Home(ctx, state)
	})

	router.GET("/about", func(ctx *gin.Context) {
		handlers.About(ctx, state)
	})

	router.GET("/projects", func(ctx *gin.Context) {
		handlers.Projects(ctx, state)
	})

	router.GET("/experience", func(ctx *gin.Context) {
		handlers.Experience(ctx, state)
	})

	// Load templates
	router.LoadHTMLGlob("views/**/*.html")

	// Static files
	router.Static("/static", "./static")

	return router
}
