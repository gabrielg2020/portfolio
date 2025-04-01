package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gabrielg2020/backend/queries"
	"github.com/gabrielg2020/backend/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func GetExperiences(c *gin.Context) {
	db := utils.OpenDatabase()
	defer db.Close()

	var experiences []map[string]interface{}

	// Get Experince query
	query := queries.GetExperiences(4)

	rows, err := db.Query(query)

	if err != nil {
		log.Fatalf("Failed to query experiences: %v", err)
	}
	defer rows.Close()

	// Grab experiences
	for rows.Next() {
		var experienceID, organisation, role, startYear, endYear, description string
		var languagesStr, technologiesStr sql.NullString

		if err := rows.Scan(&experienceID, &organisation, &role, &startYear, &endYear, &description, &languagesStr, &technologiesStr); err != nil {
			log.Fatalf("Failed to scan project: %v", err)
		}

		// Parse string 'arrays'
		var languages, technologies []string
		if languagesStr.Valid {
			languages = utils.ParseStringSlice(languagesStr.String)
		}
		if technologiesStr.Valid {
			technologies = utils.ParseStringSlice(technologiesStr.String)
		}

		// Instert data into JSON format
		experience := map[string]interface{}{
			"id":           experienceID,
			"organisation": organisation,
			"role":         role,
			"startYear":    startYear,
			"endYear":      endYear,
			"description":  description,
			"languages":    languages,
			"technologies": technologies,
		}

		experiences = append(experiences, experience)

		if err := rows.Err(); err != nil {
			log.Fatalf("Error iterating over rows: %v", err)
		}
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"experiences": experiences,
	})
}
