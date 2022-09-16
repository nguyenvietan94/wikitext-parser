package tree

type Elem interface {
	GetPlainText() (string, error)
}

type Wikicode struct {
	list []Elem
}

func NewWikicode() *Wikicode {
	w := &Wikicode{}
	return w
}

func (w *Wikicode) GetPlainText() (string, error) {
	var out string
	for _, elem := range w.list {
		next, err := elem.GetPlainText()
		if err != nil {
			return "", err
		}
		out += next
	}
	return out, nil
}

func (w *Wikicode) AddItem(e Elem) {
	w.list = append(w.list, e)
}

func (w *Wikicode) GetElemList() []Elem {
	return w.list
}
