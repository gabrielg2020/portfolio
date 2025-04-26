package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/state"
	"github.com/gin-gonic/gin"
)

func About(ctx *gin.Context, state *state.State) {
	state.SetCurrentLinkIndexByHref("/about")
	link := state.GetCurrentLink()
	ctx.HTML(http.StatusOK, "about.html", gin.H{
		"title":    link.Name,
		"prevLink": state.Previous(),
		"nextLink": state.Next(),
	})
}
