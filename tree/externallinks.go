package tree

// eg. [https://www.wikipedia.org Wikipedia]
type ExternalLink struct {
	link, displayedText string
}

func NewExternalLink(link, displayedText string) *ExternalLink {
	e := new(ExternalLink)
	e.link = link
	e.displayedText = displayedText
	return e
}

func (e *ExternalLink) GetPlainText() (string, error) {
	if e == nil {
		return "", nil
	}
	return e.link, nil
}
