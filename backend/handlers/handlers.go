package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gabrielg2020/backend/queries"
	"github.com/gabrielg2020/backend/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func HelloHandler(c *gin.Context) {
	db := openDatabase()
	defer db.Close()

	start := time.Now()
	err := db.Ping()
	elapsed := time.Since(start)

	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Printf("Successfully pinged SQLite database in %v\n", elapsed)

	c.JSON(http.StatusOK, gin.H{
		"message": "Connected to database",
	})
}

func GetProjects(c *gin.Context) {
	db := openDatabase()
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
		fmt.Print(project)

		if err := rows.Err(); err != nil {
			log.Fatalf("Error iterating over rows: %v", err)
		}
	}

	// Return Data
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/portfolio.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db.SetConnMaxLifetime(5 * time.Second)

	return db
}
