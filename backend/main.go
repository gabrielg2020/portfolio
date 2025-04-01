package main

import (
	"os"
	"github.com/gabrielg2020/backend/logger"

	"github.com/gabrielg2020/backend/routes"
)

func main() {
	// If not in development set a log level
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		logger.SetLogLevel("Warn")
	}
	
	router := routes.SetupRouter()

	logger.Info("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}
