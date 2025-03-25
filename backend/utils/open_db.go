package utils

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/portfolio.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db.SetConnMaxLifetime(5 * time.Second)

	return db
}
