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
	db, err := sql.Open("sqlite3", "./database/portfolio.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(5 * time.Second)

	start := time.Now()
	err = db.Ping()
	elapsed := time.Since(start)

	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Printf("Successfully pinged SQLite database in %v\n", elapsed)

	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}