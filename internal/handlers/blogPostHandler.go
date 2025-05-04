package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func BlogPost(ctx *gin.Context, navigation *models.Navigation, blog *models.Blog) {
	navigation.SetCurrentLinkIndexByHref("/blog")
	link := navigation.GetCurrentLink()
	ctx.HTML(http.StatusOK, blog.HtmlFile, gin.H{
		"title":    link.Name,
		"prevLink": navigation.Previous(),
		"nextLink": navigation.Next(),
		"blog":     blog,
	})
}
