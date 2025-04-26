package main

import (
	"fmt"

	"github.com/gabrielg2020/portfolio/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		fmt.Println("uh oh")
	}
}
