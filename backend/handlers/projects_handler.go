package handlers

import (
	"net/http"

	"github.com/gabrielg2020/backend/models"
	"github.com/gabrielg2020/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project

	err := utils.UnmarshalJSON("./data/projects.json", &projects)

	if err != nil {
		handleInternalServerError(c, err, "Failed to unmasall JSON: %v")
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}
