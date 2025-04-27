package models

type State struct {
	links            []Link
	currentLinkIndex int
}

func NewState() *State {
	links := []Link{
		{Name: "00 - Home", Href: "/"},
		{Name: "01 - About Me", Href: "/about"},
		{Name: "02 - Projects", Href: "/projects"},
		{Name: "03 - Experience", Href: "/experience"},
		{Name: "04 - Blog", Href: "/blog"},
	}
	return &State{
		links:            links,
		currentLinkIndex: 0,
	}
}

func (s *State) Current() *Link {
	if s.currentLinkIndex >= 0 && s.currentLinkIndex < len(s.links) {
		return &s.links[s.currentLinkIndex]
	}
	return nil
}

func (s *State) Next() *Link {
	if s.currentLinkIndex+1 < len(s.links) {
		return &s.links[s.currentLinkIndex+1]
	}
	return nil
}

func (s *State) Previous() *Link {
	if s.currentLinkIndex-1 >= 0 {
		return &s.links[s.currentLinkIndex-1]
	}
	return nil
}

func (s *State) SetCurrentLinkIndexByHref(href string) {
	for i, link := range s.links {
		if link.Href == href {
			s.currentLinkIndex = i
			return
		}
	}
}

func (s *State) GetCurrentLink() *Link {
	return &s.links[s.currentLinkIndex]
}

func (s *State) GetAllLinks() []Link {
	return s.links
}
