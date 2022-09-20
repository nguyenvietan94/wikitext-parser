package tokenizer

import (
	"regexp"
	"strings"
)

var (
	// notes:
	// {{{, }}} are ahead of {{, }};
	// ''' is ahead of ''
	/// <ref> is head of <ref
	delimiters = []string{
		TemplateParamOpen, TemplateParamClose,
		TemplateOpen, TemplateClose,
		WikilinkOpen, WikilinkClose,
		TemplateParamSeparator,
		TemplateParamEquals,
		TemplateAsteriskInList,
		TagRefOpen, TagRefClose,
		TagRefOpen1, TagClose,
		TagBreak, TagBreak1,
		TagListItem,
		TagGreaterThan, TagLessThan,
		CommentStart, CommentEnd,
		Bold, Italic,
		TextWithEscapes,
	}
	delimiterNames = map[string]string{
		"{{":     "templateOpen",
		"}}":     "templateClose",
		"[[":     "wikilinkOpen",
		"]]":     "wikilinkClose",
		"{{{":    "parameterOpen",
		"}}}":    "parameterClose",
		"<ref>":  "tagRefOpen",
		"<ref":   "tagRefOpen",
		"</ref>": "tagRefClose",
		"<br>":   "break",
		"<br />": "break1",
		"/>":     "tagClose",
		"<!--":   "commentStart",
		"-->":    "commentEnd",
		">":      "tagGreaterThan",
		"<":      "tagLessThan",
		"<li>":   "tagListItem",
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
	tokens = fixEndReferenceTag(tokens)
	return tokens, nil
}

func TokenizeWithFormatCorrection(data string) ([]*Token, error) {
	tokens, err := Tokenize(data)
	if err != nil {
		return nil, err
	}
	addSpaceBeforePattern(tokens, "(")
	addSpaceAfterPattern(tokens, ")")
	addSpaceAfterPattern(tokens, ",")
	return tokens, nil
}

// <ref https://hello.com /> => <ref, https://hello.com/, > => <ref, https://hello.com, />
func fixEndReferenceTag(tokens []*Token) []*Token {
	var out []*Token
	for i := 0; i < len(tokens); i++ {
		if i+1 < len(tokens) && strings.HasSuffix(tokens[i].Token, "/") && tokens[i+1].Token == ">" {
			first := tokens[i].Token[:len(tokens[i].Token)-1] // remove '/' at the end
			second := "/>"
			out = append(out, &Token{"text", first})
			out = append(out, &Token{"tagClose", second})
			i++
		} else {
			out = append(out, tokens[i])
		}
	}
	return out
}

func addSpaceBeforePattern(tokens []*Token, pattern string) {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Token == pattern {
			tokens[i].Token = " " + pattern
		}
	}
}

func addSpaceAfterPattern(tokens []*Token, pattern string) {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Token == pattern {
			tokens[i].Token = pattern + " "
		}
	}
}
