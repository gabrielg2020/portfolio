package main

import (
	"fmt"

	"github.com/gabrielg2020/portfolio/internal/models"
	"github.com/gabrielg2020/portfolio/internal/routes"
)

func main() {
	// Initilise State
	state := models.NewState()

	router := routes.SetupRouter(state)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("uh oh")
	}
}
