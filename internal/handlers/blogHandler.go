package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Blog(ctx *gin.Context, navigation *models.Navigation) {
	navigation.SetCurrentLinkIndexByHref("/blog")
	link := navigation.GetCurrentLink()
	ctx.HTML(http.StatusOK, "blog.html", gin.H{
		"title":    link.Name,
		"prevLink": navigation.Previous(),
		"nextLink": navigation.Next(),
	})
}
