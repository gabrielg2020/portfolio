package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context, state *models.State) {
	state.SetCurrentLinkIndexByHref("/")
	link := state.GetCurrentLink()
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title":    link.Name,
		"prevLink": state.Previous(),
		"nextLink": state.Next(),
		"links":    state.GetAllLinks()[1:],
	})
}
