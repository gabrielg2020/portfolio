package main

import (
	"github.com/gabrielg2020/portfolio/backend/routes"
	"log"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server starting on :8080")
	router.Run(":8080")
}