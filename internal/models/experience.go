package models

type Experience struct {
	Type         string
	Organisation string
	Role         string
	Description  string
	Languages    string
	Technologies string
	StartYear    string
	EndYear      string
}

func NewExperience(name, organisation, role, descriprtion, languages, technologies, startYear, endYear string) *Experience {
	return &Experience{
		Type:         name,
		Organisation: organisation,
		Role:         role,
		Description:  descriprtion,
		Languages:    languages,
		Technologies: technologies,
		StartYear:    startYear,
		EndYear:      endYear,
	}
}
