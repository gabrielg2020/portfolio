package handlers

import (
	"net/http"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gin-gonic/gin"
)

func Projects(ctx *gin.Context, state *models.State) {
	projects := getProjects()

	state.SetCurrentLinkIndexByHref("/projects")
	link := state.GetCurrentLink()
	ctx.HTML(http.StatusOK, "projects.html", gin.H{
		"title":    link.Name,
		"prevLink": state.Previous(),
		"nextLink": state.Next(),
		"projects": projects,
	})
}

func getProjects() []models.Project {
	return []models.Project{
		*models.NewProject(
			"Portfolio",
			"Portfolio Description",
			"Maintaing",
			"02/04/2025",
			"Go",
			"Docker",
			"https://www.github.com/gabrielg2020/portfolio",
			"https://www.gabrielg.tech",
		),
		*models.NewProject(
			"Maze Visualiser",
			"Maze Visualiser Description",
			"Built",
			"22/02/2023",
			"VB.Net",
			"",
			"https://www.github.com/gabrielg2020/maze-visualiser",
			"",
		),
	}
}
