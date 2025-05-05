package models

type Project struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	GithubURL    string   `json:"githubURL"`
	LiveURL      string   `json:"liveURL"`
	Languages    []string `json:"languages"`
	Technologies []string `json:"technologies"`
}

