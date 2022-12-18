package tree

import (
	"strings"
	"wikitext-parser/utils"
)

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
			continue
		}
		if len(out) == 0 && utils.IsPuctuation(next) {
			continue
		}
		if strings.HasSuffix(out, ",") && next == "," {
			continue
		}
		if len(out) > 0 && !strings.HasPrefix(next, ",") && !strings.HasPrefix(next, "(") && !strings.HasPrefix(next, ")") && !strings.HasSuffix(out, "(") {
			out += " "
		}
		out += next
	}
	return out, nil
}

func (w *Wikicode) AddItem(e Elem) {
	if e != nil {
		w.list = append(w.list, e)
	}
}

func (w *Wikicode) GetElemList() []Elem {
	return w.list
}
