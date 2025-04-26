package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	links := []models.Link{
		{Name: "01 - About Me", Href: "/about"},
		{Name: "02 - Projects", Href: "/projects"},
		{Name: "03 - Experience", Href: "/experience"},
		{Name: "04 - Blog", Href: "/blog"},
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
		"links": links,
	})
}
