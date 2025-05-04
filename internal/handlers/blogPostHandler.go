package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gabrielg2020/portfolio/internal/services"
	"github.com/gin-gonic/gin"
)

func BlogPost(ctx *gin.Context, navigation *models.Navigation, blog *models.Blog) {
	navigation.SetCurrentLinkIndexByHref("/blog")
	ctx.HTML(http.StatusOK, "blogLayout.html", gin.H{
		"prevLink": navigation.Previous(),
		"nextLink": navigation.Next(),
		"blog":     blog,
		"content":  services.LoadHTML("/home/gabriel/projects/portfolio/views/blog/test.html"),
	})
}
