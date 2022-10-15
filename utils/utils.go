package utils

import (
	"html"
	"strings"
)

var (
	infoboxPatterns     = []string{"{{Infobox", "{{Tóm tắt", "{{Thông tin", "{{Hộp thông tin"}
	infoboxNamePrefixes = []string{"hộp thông tin", "thông tin", "infobox"}
)

const MAX = 1000000000

func PreprocessText(text string) string {
	text = html.UnescapeString(text)
	text = strings.ReplaceAll(text, "\u00a0", " ")
	return text
}

func PreprocessTemplateName(templateName string) string {
	templateName = strings.ToLower(templateName)
	for _, prefix := range infoboxNamePrefixes {
		templateName = strings.Replace(templateName, prefix, "", 1)
	}
	return strings.TrimSpace(templateName)
}

func IsPuctuation(text string) bool {
	return text == "." || text == ","
}

// get wiki infobox from wikipedia text
func GetWikiInfoboxFromText(text string) string {
	start := MAX
	for _, pattern := range infoboxPatterns {
		idx := strings.Index(text, pattern)
		if idx >= 0 {
			if start == MAX {
				start = idx
			} else if idx < start {
				start = idx
			}
		}
	}
	if start < MAX {
		cntCurlyBrackets := 2
		for i := start + 2; i < len(text); i++ {
			if text[i] == '{' {
				cntCurlyBrackets++
			} else if text[i] == '}' {
				cntCurlyBrackets--
			}
			if cntCurlyBrackets == 0 {
				return text[start : i+1]
			}
		}
	}
	return ""
}
