package main

import (
	"log"

	"github.com/gabrielg2020/backend/routes"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server starting on :8080")
	router.Run(":8080")
}
