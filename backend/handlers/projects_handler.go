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

func GetProjects(c *gin.Context) {
	db := utils.OpenDatabase()
	defer db.Close()

	var projects []map[string]interface{}

	// Get Project query
	query := queries.GetProjects(4)

	rows, err := db.Query(query)

	if err != nil {
		log.Fatalf("Failed to query projects: %v", err)
	}
	defer rows.Close()

	// Grab projects
	for rows.Next() {
		var projectID, title, description, githubLink string
		var languagesStr, technologiesStr sql.NullString

		if err := rows.Scan(&projectID, &title, &description, &githubLink, &languagesStr, &technologiesStr); err != nil {
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
		project := map[string]interface{}{
			"id":           projectID,
			"title":        title,
			"description":  description,
			"githubLink":   githubLink,
			"languages":    languages,
			"technologies": technologies,
		}

		projects = append(projects, project)

		if err := rows.Err(); err != nil {
			log.Fatalf("Error iterating over rows: %v", err)
		}
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}
