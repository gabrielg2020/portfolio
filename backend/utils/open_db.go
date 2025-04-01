package utils

import (
	"database/sql"
	"time"

	"github.com/gabrielg2020/backend/logger"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDatabase returns a sql.DB struct with connection to porfolio.db.
//
// Example:
//
//	db := OpenDatabase()
func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/portfolio.db")
	if err != nil {
		logger.Errorf("Failed to open database: %v", err)
	}

	db.SetConnMaxLifetime(5 * time.Second)

	return db
}
