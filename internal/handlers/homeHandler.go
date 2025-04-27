package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context, navigation *models.Navigation) {
	navigation.SetCurrentLinkIndexByHref("/")
	link := navigation.GetCurrentLink()
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title":    link.Name,
		"prevLink": navigation.Previous(),
		"nextLink": navigation.Next(),
		"links":    navigation.GetAllLinks()[1:],
	})
}
