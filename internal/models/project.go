package models

type Project struct {
	Name         string
	Description  string
	Status       string
	Date         string
	Languages    string
	Technologies string
	GithubURL    string
	LiveURL      string
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
