package main

import (
	"fmt"

	"github.com/gabrielg2020/portfolio/internal/routes"
	"github.com/gabrielg2020/portfolio/internal/state"
)

func main() {
	// Initilise State
	state := state.NewState()

	router := routes.SetupRouter(state)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("uh oh")
	}
}
