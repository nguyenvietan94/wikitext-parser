package tree

import "wikitext-parser/tokenizer"

type Text string

func NewText(data string) Text {
	t := Text(data)
	return t
}

func (t Text) GetPlainText() (string, error) {
	out := string(t)
	if out == tokenizer.TemplateAsteriskInList { // * in list
		out = ", "
	}
	return out, nil
}
