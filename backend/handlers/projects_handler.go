package handlers

import (
	"database/sql"
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
		handleInternalServerError(c, err, "Failed to query projects: %v")
	}
	defer rows.Close()

	// Grab projects
	for rows.Next() {
		var projectID, title, description, githubURL string
		var languagesStr, technologiesStr, liveURLStr sql.NullString

		if err := rows.Scan(&projectID, &title, &description, &githubURL, &liveURLStr, &languagesStr, &technologiesStr); err != nil {
			handleInternalServerError(c, err, "Failed to scan project: %v")
		}

		// Parse string 'arrays'
		var languages, technologies []string
		if languagesStr.Valid {
			languages = utils.ParseStringSlice(languagesStr.String)
		}
		if technologiesStr.Valid {
			technologies = utils.ParseStringSlice(technologiesStr.String)
		}

		// Check if liveURL is valid
		var liveURL interface{}
		if liveURLStr.Valid {
			liveURL = liveURLStr.String
		} else {
			liveURL = nil
		}

		// Instert data into JSON format
		project := map[string]interface{}{
			"id":           projectID,
			"title":        title,
			"description":  description,
			"githubURL":    githubURL,
			"liveURL":      liveURL,
			"languages":    languages,
			"technologies": technologies,
		}

		projects = append(projects, project)

		if err := rows.Err(); err != nil {
			handleInternalServerError(c, err, "Error iterating over rows: %v")
		}
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}
