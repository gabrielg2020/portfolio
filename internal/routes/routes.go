package routes

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/handlers"
	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter(navigation *models.Navigation, storage *models.Storage) *gin.Engine {
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
		handlers.Home(ctx, navigation)
	})

	router.GET("/about", func(ctx *gin.Context) {
		handlers.About(ctx, navigation)
	})

	router.GET("/projects", func(ctx *gin.Context) {
		handlers.Projects(ctx, navigation, storage)
	})

	router.GET("/experience", func(ctx *gin.Context) {
		handlers.Experience(ctx, navigation, storage)
	})

	router.GET("/blog", func(ctx *gin.Context) {
		handlers.Blog(ctx, navigation, storage)
	})

	blog := router.Group("/blog")
	{
		for _, blogPost := range storage.Blogs {
			blog.GET(blogPost.Href, func(ctx *gin.Context) {
				handlers.BlogPost(ctx, navigation, &blogPost)
			})
		}
	}

	// Load templates
	router.LoadHTMLGlob("views/**/*.html")

	// Static files
	router.Static("/static", "./static")

	return router
}
