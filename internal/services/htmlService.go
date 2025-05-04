package services

import (
	"html/template"
	"log"
	"os"
)

func LoadHTML(htmlPath string) template.HTML {
	htmlContent, err := os.ReadFile(htmlPath)
	if err != nil {
		log.Printf("Warning: could not read html file")
	}

	return template.HTML(htmlContent)
}
