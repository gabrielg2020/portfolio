package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Experience(ctx *gin.Context, navigation *models.Navigation, storage *models.Storage) {
	navigation.SetCurrentLinkIndexByHref("/experience")
	link := navigation.GetCurrentLink()
	ctx.HTML(http.StatusOK, "experiences.html", gin.H{
		"title":       link.Name,
		"experiences": storage.Experiences,
	})
}
