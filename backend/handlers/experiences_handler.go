package handlers

import (
	"net/http"

	"github.com/gabrielg2020/backend/models"
	"github.com/gabrielg2020/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetExperiences(c *gin.Context) {
	var experiences []models.Experience

	err := utils.UnmarshalJSON("./data/experiences.json", &experiences)

	if err != nil {
		handleInternalServerError(c, err, "Failed to unmasall JSON: %v")
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"experiences": experiences,
	})
}
