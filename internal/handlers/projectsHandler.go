package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Projects(ctx *gin.Context, navigation *models.Navigation, storage *models.Storage) {
	navigation.SetCurrentLinkIndexByHref("/projects")
	link := navigation.GetCurrentLink()
	ctx.HTML(http.StatusOK, "projects.html", gin.H{
		"title":    link.Name,
		"projects": storage.Projects,
	})
}
