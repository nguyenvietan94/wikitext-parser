package parser

import (
	"fmt"
	"strconv"
	"wikitext-parser/templates"
	"wikitext-parser/tokenizer"
	"wikitext-parser/tree"
	"wikitext-parser/utils"
)

type fn func() tree.Elem

type fieldValue struct {
	id    int
	value string
}

type Parser struct {
	data      string
	tokens    []*tokenizer.Token
	cTokenIdx int
	root      *tree.Wikicode
	handlers  map[string]fn
	params    map[string]*fieldValue
}

func NewParser(data string) *Parser {
	p := &Parser{data: utils.PreprocessText(data)}
	p.root = tree.NewWikicode()
	p.handlers = map[string]fn{
		"templateOpen":  p.handleTemplate,
		"templateClose": p.handleTemplate,
		"text":          p.handleText,
		"wikilinkOpen":  p.handleWikilink,
		"wikilinkClose": p.handleWikilink,
		"tagRefOpen":    p.handleReference,
		"tagRefClose":   p.handleReference,
		"commentStart":  p.handleComment,
		"commentEnd":    p.handleComment,
		"break":         p.handleTagBreak,
	}
	p.params = make(map[string]*fieldValue)
	return p
}

func (p *Parser) parse() error {
	err := p.tokenize()
	if err != nil {
		return err
	}
	for ; p.cTokenIdx < len(p.tokens); p.cTokenIdx++ {
		tokenName := p.tokens[p.cTokenIdx].Type
		if f, ok := p.handlers[tokenName]; ok {
			item := f()
			if item != nil {
				p.root.AddItem(item)
			}
		}
	}
	err = p.processParams()
	return err
}

func (p *Parser) tokenize() error {
	var err error
	p.tokens, err = tokenizer.Tokenize(p.data)
	return err
}

// TODO: read template format from templates.json
func (p *Parser) processParams() error {
	for _, elem := range p.root.GetElemList() {
		if template, ok := elem.(*tree.Template); ok {
			if len(template.Name) > 0 {
				expectedFields := templates.GetFieldsFromTemplate(template.Name)
				for _, field := range expectedFields {
					if text, err := template.GetPlainTextByField(field.En); err == nil {
						p.params[field.DisplayedText] = &fieldValue{id: field.Id, value: text}
					}
					if text, err := template.GetPlainTextByField(field.Vi); err == nil {
						p.params[field.DisplayedText] = &fieldValue{id: field.Id, value: text}
					}
				}
			}
		}
	}
	return nil
}

func (p *Parser) getParams() map[string]*fieldValue {
	return p.params
}

func (p *Parser) printTokens() {
	for _, token := range p.tokens {
		fmt.Println(token.Token)
	}
}

// --

func (p *Parser) handleTemplate() tree.Elem {
	template := tree.NewTemplate()
	var token, prevToken *tokenizer.Token
	defaultIdx := 1
	for ; p.cTokenIdx < len(p.tokens); p.cTokenIdx++ {
		token = p.tokens[p.cTokenIdx]
		if token.Token == tokenizer.TemplateClose { // }}, end of the template
			break
		}
		if token == nil || token.Token == tokenizer.TemplateOpen { // {{
			continue
		}
		if p.cTokenIdx > 0 {
			prevToken = p.tokens[p.cTokenIdx-1]
		}
		if prevToken != nil && prevToken.Token == tokenizer.TemplateOpen { // {{ name
			template.Name = utils.PreprocessTemplateName(token.Token)
			continue
		}
		if token.Token == tokenizer.TemplateParamSeparator || token.Token == tokenizer.TemplateAsteriskInList { // |, *
			if p.cTokenIdx+2 < len(p.tokens) && p.tokens[p.cTokenIdx+2].Token == tokenizer.TemplateParamEquals { // =
				p.cTokenIdx += 3
				key := p.tokens[p.cTokenIdx-2].Token
				template.Params[key] = p.handleTemplateParams()
			} else {
				if p.cTokenIdx+1 < len(p.tokens) {
					p.cTokenIdx++
					template.Params[strconv.Itoa(defaultIdx)] = p.handleTemplateParams()
					defaultIdx++
				}
			}
		}
	}
	return template
}

func (p *Parser) handleTemplateParams() *tree.Wikicode {
	wikicode := tree.NewWikicode()
	for ; p.cTokenIdx < len(p.tokens); p.cTokenIdx++ {
		if p.tokens[p.cTokenIdx].Token == tokenizer.TemplateParamSeparator || p.tokens[p.cTokenIdx].Token == tokenizer.TemplateClose { // |
			p.cTokenIdx--
			break
		}
		tokenType := p.tokens[p.cTokenIdx].Type
		if f, ok := p.handlers[tokenType]; ok {
			item := f()
			if item != nil {
				wikicode.AddItem(item)
			}
		}
	}
	return wikicode
}

// skip reference tags
func (p *Parser) handleReference() tree.Elem {
	for ; p.cTokenIdx < len(p.tokens); p.cTokenIdx++ {
		if p.tokens[p.cTokenIdx].Token == tokenizer.TagRefClose {
			break
		}
	}
	return nil
}

func (p *Parser) handleText() tree.Elem {
	text := tree.NewText(p.tokens[p.cTokenIdx].Token)
	return text
}

func (p *Parser) handleWikilink() tree.Elem {
	var wikilink *tree.Wikilink
	if p.tokens[p.cTokenIdx].Token == tokenizer.WikilinkOpen { // [[
		if p.cTokenIdx+2 < len(p.tokens) && p.tokens[p.cTokenIdx+2].Token == tokenizer.WikilinkClose { // [[wikiPage]]
			text := p.tokens[p.cTokenIdx+1].Token
			p.cTokenIdx += 2
			wikilink = tree.NewWikilink(text, "")
		} else if p.cTokenIdx+4 < len(p.tokens) && p.tokens[p.cTokenIdx+4].Token == tokenizer.WikilinkClose { // [[wikiPage|displayedText]]
			if p.cTokenIdx+3 < len(p.tokens) {
				wikiPage := p.tokens[p.cTokenIdx+1].Token
				displayedText := p.tokens[p.cTokenIdx+3].Token
				p.cTokenIdx += 4
				wikilink = tree.NewWikilink(wikiPage, displayedText)
			}
		} else {
			if p.cTokenIdx+1 < len(p.tokens) {
				wikiPage := p.tokens[p.cTokenIdx+1].Token
				wikilink = tree.NewWikilink(wikiPage, "")
			}
			for ; p.cTokenIdx < len(p.tokens) && p.tokens[p.cTokenIdx].Token != tokenizer.WikilinkClose; p.cTokenIdx++ {
			}
		}
	}
	return wikilink
}

// skip comment tags
func (p *Parser) handleComment() tree.Elem {
	for ; p.cTokenIdx < len(p.tokens); p.cTokenIdx++ {
		if p.tokens[p.cTokenIdx].Token == tokenizer.CommentEnd {
			break
		}
	}
	return nil
}

func (p *Parser) handleTagBreak() tree.Elem {
	return nil
}
