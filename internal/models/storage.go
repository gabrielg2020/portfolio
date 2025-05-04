package models

type Storage struct {
	Projects    []Project    `json:"projects"`
	Experiences []Experience `json:"experiences"`
	Blogs []Blog `json:"blogs"`
}
