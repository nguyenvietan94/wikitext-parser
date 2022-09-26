package parser

import (
	"fmt"
	"strconv"
	"strings"
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
	params    map[string]string
}

func NewParser(data string) *Parser {
	p := &Parser{data: utils.PreprocessText(data)}
	p.root = tree.NewWikicode()
	p.handlers = map[string]fn{
		"templateOpen":      p.handleTemplate,
		"templateClose":     p.handleTemplate,
		"text":              p.handleText,
		"wikilinkOpen":      p.handleWikilink,
		"wikilinkClose":     p.handleWikilink,
		"externalLinkOpen":  p.handleExternalLink,
		"externalLinkClose": p.handleExternalLink,
		"tagRefOpen":        p.handleReference,
		"tagRefClose":       p.handleReference,
		"commentStart":      p.handleComment,
		"commentEnd":        p.handleComment,
		"break":             p.handleTagBreak,
	}
	p.params = make(map[string]string)
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
	err = p.convertTree2PlainText()
	return err
}

func (p *Parser) tokenize() error {
	var err error
	p.tokens, err = tokenizer.Tokenize(p.data)
	return err
}

func (p *Parser) convertTree2PlainText() error {
	elemList := p.root.GetElemList()
	if len(elemList) != 1 {
		return fmt.Errorf("only support element list of length 1, but length=%d found", len(elemList))
	}
	if template, ok := elemList[0].(*tree.Template); ok && template != nil {
		p.params = template.GetParamsInPlainText()
	}
	return nil
}

// returns fields defined in templates.json if available
func (p *Parser) getRequiredFields() (map[string]*fieldValue, error) {
	elemList := p.root.GetElemList()
	if len(elemList) != 1 {
		return nil, fmt.Errorf("only support element list of length 1, but length=%d found", len(elemList))
	}
	p.mergeParams()
	out := make(map[string]*fieldValue)
	if template, ok := elemList[0].(*tree.Template); ok && template != nil {
		if len(template.Name) > 0 {
			expectedFields := templates.GetFieldsFromTemplate(template.Name)
			for _, field := range expectedFields {
				if field.Enabled {
					if text, ok := p.params[field.En]; ok {
						out[field.DisplayedText] = &fieldValue{id: field.Id, value: text}
					}
					if text, ok := p.params[field.Vi]; ok {
						out[field.DisplayedText] = &fieldValue{id: field.Id, value: text}
					}
				}
			}
		}
	}
	return out, nil
}

func (p *Parser) merge(from, to string) {
	var textFrom string
	var ok1, ok2 bool
	textFrom, ok1 = p.params[from]
	_, ok2 = p.params[to]
	if (!ok1 && !ok2) || textFrom == "" {
		return
	}
	if strings.HasSuffix(from, "_year") {
		textFrom = "(" + textFrom + ")"
	} else {
		textFrom = ", " + textFrom
	}
	p.params[to] += textFrom
}

func (p *Parser) mergeParams() {
	p.merge("equity_year", "equity")
	p.merge("income_year", "operating_income")
	p.merge("revenue_year", "revenue")
	p.merge("assets_year", "assets")
	p.merge("num_locations_year", "num_locations")
	p.merge("num_employees_year", "num_employees")
	p.merge("net_income_year", "net_income")
	p.merge("location_country", "location_city")
	p.merge("hq_location_country", "hq_location_city")
}

func (p *Parser) getParams() map[string]string {
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
		if p.tokens[p.cTokenIdx].Token == tokenizer.TagRefClose || p.tokens[p.cTokenIdx].Token == tokenizer.TagClose {
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
	return tree.NewText(",")
}

func (p *Parser) handleExternalLink() tree.Elem {
	var externalLink *tree.ExternalLink
	if p.cTokenIdx+2 < len(p.tokens) && p.tokens[p.cTokenIdx+2].Token == tokenizer.ExternalLinkClose {
		externalLink = tree.NewExternalLink(p.tokens[p.cTokenIdx+1].Token, "")
		p.cTokenIdx += 2
	} else if p.cTokenIdx+3 < len(p.tokens) && p.tokens[p.cTokenIdx+3].Token == tokenizer.ExternalLinkClose {
		externalLink = tree.NewExternalLink(p.tokens[p.cTokenIdx+2].Token, "")
		p.cTokenIdx += 3
	}
	return externalLink
}
