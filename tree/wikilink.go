package tree

import "strings"

var (
	prefixes = []string{"File:", ":File:", "Tập tin:", ":Tập tin:", "Media:"}
)

type Wikilink struct {
	wikiPage      string
	displayedText string
}

func NewWikilink(wikiPage, displayedText string) *Wikilink {
	w := new(Wikilink)
	w.wikiPage = wikiPage
	w.displayedText = displayedText
	return w
}

func (w *Wikilink) GetPlainText() (string, error) {
	if len(w.displayedText) > 0 {
		return w.displayedText, nil
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(w.wikiPage, prefix) {
			w.wikiPage = strings.Replace(w.wikiPage, prefix, "", 1)
			break
		}
	}
	return w.wikiPage, nil
}
