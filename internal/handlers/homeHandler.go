package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/state"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context, state *state.State) {

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title": state.CurrentLink.Name,
		"links": state.Links[1:],
	})
}
