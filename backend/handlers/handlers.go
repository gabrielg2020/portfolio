package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func HelloHandler(c *gin.Context) {
	db := openDatabase()

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

	// Query database
	rows, err := db.Query("SELECT title, description, github_link FROM projects")
	if err != nil {
		log.Fatalf("Failed to query projects: %v", err)
	}
	defer rows.Close()

	var projects []map[string]interface{}

	// Grab all projects
	for rows.Next() {
		var title, description, githubLink string
		if err := rows.Scan(&title, &description, &githubLink); err != nil {
			log.Fatalf("Failed to scan project: %v", err)
		}
		project := map[string]interface{}{
			"title":       title,
			"description": description,
			"githubLink":  githubLink,
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

	// Send Response
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/portfolio.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(5 * time.Second)

	return db
}
