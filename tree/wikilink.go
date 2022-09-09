package tree

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
	return w.wikiPage, nil
}
