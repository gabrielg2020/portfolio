package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func About(ctx *gin.Context, navigation *models.Navigation) {
	navigation.SetCurrentLinkIndexByHref("/about")
	link := navigation.GetCurrentLink()
	ctx.HTML(http.StatusOK, "about.html", gin.H{
		"title":    link.Name,
		"prevLink": navigation.Previous(),
		"nextLink": navigation.Next(),
	})
}
