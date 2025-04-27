package models

type Project struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Status       string `json:"staus"`
	Date         string `json:"date"`
	Languages    string `json:"languages"`
	Technologies string `json:"technologies"`
	GithubURL    string `json:"githubURL"`
	LiveURL      string `json:"liveURL"`
}

func NewProject(name, description, status, date, languages, technologies, githubURL, liveURL string) *Project {
	return &Project{
		Name:         name,
		Description:  description,
		Status:       status,
		Date:         date,
		Languages:    languages,
		Technologies: technologies,
		GithubURL:    githubURL,
		LiveURL:      liveURL,
	}
}
