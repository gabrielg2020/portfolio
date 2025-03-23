package main

import (
	"log"

	"github.com/gabrielg2020/backend/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")

	log.Println("Server starting on :8080")
}
