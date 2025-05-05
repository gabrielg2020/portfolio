package models

type Experience struct {
	ID           int      `json:"id"`
	Type         string   `json:"type"`
	Organisation string   `json:"organisation"`
	Role         string   `json:"role"`
	StartYear    string   `json:"startYear"`
	EndYear      string   `json:"endYear"`
	Description  string   `json:"description"`
	Languages    []string `json:"languages"`
	Technologies []string `json:"technologies"`
}
