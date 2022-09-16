package utils

import (
	"html"
	"strings"
)

func PreprocessText(text string) string {
	text = html.UnescapeString(text)
	text = strings.ReplaceAll(text, "\u00a0", " ")
	return text
}

func PreprocessTemplateName(templateName string) string {
	prefixes := []string{"hộp thông tin", "thông tin", "infobox"}
	templateName = strings.ToLower(templateName)
	for _, prefix := range prefixes {
		templateName = strings.Replace(templateName, prefix, "", 1)
	}
	return strings.TrimSpace(templateName)
}
