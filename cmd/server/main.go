package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gabrielg2020/portfolio/internal/routes"
	"github.com/gabrielg2020/portfolio/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initilise navigation
	navigation := models.NewNavigation()

	// Load storage
	storage, err := services.LoadStorage(os.Getenv("CONTENT_DIR_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	router := routes.SetupRouter(navigation, storage)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("uh oh")
	}
}
