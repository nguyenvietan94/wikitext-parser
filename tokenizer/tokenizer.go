package tokenizer

import (
	"regexp"
	"strings"
)

var (
	// notes:
	// {{{, }}} are placed ahead of {{, }};
	// ''' is placed ahead of ''
	delimiters = []string{
		TemplateParamOpen, TemplateParamClose, TemplateOpen, TemplateClose, WikilinkOpen, WikilinkClose,
		TemplateParamSeparator, TemplateParamEquals, TemplateAsteriskInList, TagRefOpen, TagRefClose, TagBreak, Bold, Italic, TextWithEscapes,
	}
	delimiterNames = map[string]string{
		"{{":     "templateOpen",
		"}}":     "templateClose",
		"[[":     "wikilinkOpen",
		"]]":     "wikilinkClose",
		"{{{":    "parameterOpen",
		"}}}":    "parameterClose",
		"<ref>":  "tagRefOpen",
		"</ref>": "tagRefClose",
		"<br>":   "break",
		"<!--":   "commentStart",
		"-->":    "commentEnd",
		"|":      "templateParamSeparator",
		"=":      "templateParamEquals",
		"''":     "italic",
		"'''":    "bold",
	}
)

type Token struct {
	Type, Token string
}

func genRegExp() string {
	regExpression := ""
	for _, delimiter := range delimiters {
		if len(regExpression) > 0 {
			regExpression += "|"
		}
		delimiterWithEscapes := addEscapes(delimiter)
		regExpression += "(" + delimiterWithEscapes + ")"
	}
	regExpression = "(" + regExpression + ")"
	return regExpression
}

func addEscapes(text string) string {
	if !strings.Contains(text, `\`) {
		text = strings.ReplaceAll(text, "[", `\[`)
		text = strings.ReplaceAll(text, "]", `\]`)
		text = strings.ReplaceAll(text, "|", `\|`)
		text = strings.ReplaceAll(text, "*", `\*`)
	}
	return text
}

func Tokenize(data string) ([]*Token, error) {
	regExpression := genRegExp()
	r, err := regexp.Compile(regExpression)
	if err != nil {
		return nil, err
	}
	parts := r.FindAllString(data, -1)

	var tokens []*Token
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			tokenName := "text"
			if val, ok := delimiterNames[p]; ok {
				tokenName = val
			}
			tokens = append(tokens, &Token{tokenName, p})
		}
	}
	return tokens, nil
}
