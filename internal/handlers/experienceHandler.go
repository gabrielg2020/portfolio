package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gabrielg2020/portfolio/internal/state"
	"github.com/gin-gonic/gin"
)

func Experience(ctx *gin.Context, state *state.State) {
	experiences := GetExperiences()

	state.SetCurrentLinkIndexByHref("/experience")
	link := state.GetCurrentLink()
	ctx.HTML(http.StatusOK, "experiences.html", gin.H{
		"title":       link.Name,
		"prevLink":    state.Previous(),
		"nextLink":    state.Next(),
		"experiences": experiences,
	})
}

func GetExperiences() []models.Experience {
	return []models.Experience{
		*models.NewExperience(
			"Work",
			"DVSA",
			"Software Engineer",
			"Work Description",
			"PHP, TypeScript",
			"Docker, Laminas, AWS, MySQL",
			"2024",
			"Current",
		),
		*models.NewExperience(
			"Education",
			"University of Roehampton",
			"BSc Digital Technology Solutions",
			"Education Description",
			"Python, Go",
			"AWS, MySQL",
			"2024",
			"Current",
		),
	}
}
