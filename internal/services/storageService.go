package services

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gabrielg2020/portfolio/internal/models"
)

func LoadStorage(dataDir string) (*models.Storage, error) {
	storage := &models.Storage{}

	// Load projects
	projects, err := loadJSONFile[[]models.Project](filepath.Join(dataDir, "projects.json"))
	if err != nil {
		return nil, err
	}
	storage.Projects = projects

	// Load experiences
	experiences, err := loadJSONFile[[]models.Experience](filepath.Join(dataDir, "experiences.json"))
	if err != nil {
		return nil, err
	}
	storage.Experiences = experiences

	// Load blogs
	blogs, err := loadJSONFile[[]models.Blog](filepath.Join(dataDir, "blogs.json"))
	if err != nil {
		return nil, err
	}
	storage.Blogs = blogs

	return storage, nil
}

func loadJSONFile[T any](filepath string) (T, error) {
	var result T

	file, err := os.ReadFile(filepath)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(file, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
