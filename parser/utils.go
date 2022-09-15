package parser

import (
	"html"
	"strings"
)

func preprocessText(text string) string {
	text = html.UnescapeString(text)
	text = strings.ReplaceAll(text, "\u00a0", " ")
	return text
}
