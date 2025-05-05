package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gabrielg2020/backend/models"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project

	// Read JSON file
	jsonContent, err := os.ReadFile("./data/projects.json")
	if err != nil {
		handleInternalServerError(c, err, "Failed to read projects.json: %v")
	}

	// Unmarshall JSON into projects
	err = json.Unmarshal(jsonContent, &projects)
	if err != nil {
		handleInternalServerError(c, err, "Failed to unmashall projects.json: %v")
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}
