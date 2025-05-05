package models

type Navigation struct {
	links            []Link
	currentLinkIndex int
}

func NewNavigation() *Navigation {
	links := []Link{
		{Name: "00 - Home", Href: "/"},
		{Name: "01 - About Me", Href: "/about"},
		{Name: "02 - Projects", Href: "/projects"},
		{Name: "03 - Experience", Href: "/experience"},
		{Name: "04 - Blog", Href: "/blog"},
	}
	return &Navigation{
		links:            links,
		currentLinkIndex: 0,
	}
}

func (n *Navigation) Current() *Link {
	if n.currentLinkIndex >= 0 && n.currentLinkIndex < len(n.links) {
		return &n.links[n.currentLinkIndex]
	}
	return nil
}

func (n *Navigation) SetCurrentLinkIndexByHref(href string) {
	for i, link := range n.links {
		if link.Href == href {
			n.currentLinkIndex = i
			return
		}
	}
}

func (n *Navigation) GetCurrentLink() *Link {
	return &n.links[n.currentLinkIndex]
}

func (n *Navigation) GetAllLinks() []Link {
	return n.links
}
