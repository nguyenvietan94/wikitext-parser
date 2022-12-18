package tree

import (
	"strings"
	"wikitext-parser/tokenizer"
)

type Text string

func NewText(data string) Text {
	t := Text(data)
	return t
}

func (t Text) GetPlainText() (string, error) {
	out := string(t)
	if out == tokenizer.TemplateAsteriskInList || out == tokenizer.TagBreak || out == tokenizer.TagBreak1 { // * in list
		out = ","
	}
	for _, prefix := range prefixes {
		out = strings.TrimPrefix(out, prefix)
	}
	return out, nil
}
