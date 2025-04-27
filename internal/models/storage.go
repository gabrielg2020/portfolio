package models

type Storage struct {
	Projects    []Project    `json:"projects"`
	Experiences []Experience `json:"experiences"`
	// BlogPosts []BlogPost `json:"blogPosts"`
}
