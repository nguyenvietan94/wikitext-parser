package tree

import (
	"fmt"
	"strings"
)

var (
	units = map[string]string{
		"metres":      "m",
		"meters":      "m",
		"m":           "m",
		"cm":          "cm",
		"centimetres": "cm",
		"centimeter":  "cm",
	}
)

type fn func() (string, error)

type Template struct {
	Name             string
	Params           map[string]*Wikicode
	templateHandlers map[string]fn
}

func NewTemplate() *Template {
	t := new(Template)
	t.Params = make(map[string]*Wikicode)
	t.templateHandlers = map[string]fn{
		// Birth date
		"birth date and age": t.handleBirthDateAndAge0,
		"birth-date and age": t.handleBirthDateAndAge1,
		"birth year and age": t.handleBirthYearAndAge,
		"ngày sinh":          t.handleBirthDateAndAge1,
		"ngày sinh và tuổi":  t.handleBirthDateAndAge1,
		"năm sinh và tuổi":   t.handleBirthYearAndAge,
		"ngày mất":           nil,
		"ngày mất và tuổi":   nil,
		"năm mất và tuổi":    nil,

		// List
		"hlist":           t.handleHList,
		"plainlist":       t.handlePlainList,
		"plain list":      t.handlePlainList,
		"flatlist":        t.handleFlatList,
		"flat list":       t.handleFlatList,
		"unbulleted list": t.handleUnbulletedList,
		"ordered list":    t.handleOrderedList,
		"pagelist":        t.handlePageList,

		// URL
		"url": t.handleURL,

		// height
		"height": t.handleHeight,

		// nowrap
		"nowrap": t.handleNowrap,
	}
	return t
}

func (t *Template) GetPlainText() (string, error) {
	key := strings.ToLower(t.Name)
	if f, ok := t.templateHandlers[key]; ok && f != nil {
		return f()
	}
	return "", fmt.Errorf("not supported template: %s", t.Name)
}

// ----- Birthdate -----

// {{Birth-date and age|1941}} → 1941 (age 81)
func (t *Template) handleBirthDateAndAge1() (string, error) {
	return t.getParamTextByKeyIndex(1)
}

// {{Birth date and age|2016|12|31|df=y}} → 31 December 2016 (age 5)
func (t *Template) handleBirthDateAndAge0() (string, error) {
	var yymmdd [3]string
	var err error
	for i := 0; i < 3; i++ {
		if wikicode, ok := t.Params[fmt.Sprintf("%d", i+1)]; ok && wikicode != nil {
			yymmdd[i], err = wikicode.GetPlainText()
			if err != nil || len(yymmdd[i]) == 0 {
				return "", err
			}
		}
	}
	return yymmdd[2] + "-" + yymmdd[1] + "-" + yymmdd[0], nil
}

// {{Birth year and age|1941}} → 1941 (age 80–81)
func (t *Template) handleBirthYearAndAge() (string, error) {
	return t.getParamTextByKeyIndex(1)
}

// ----- Lists -----

func (t *Template) handleList() (string, error) {
	var out string
	for _, wikicode := range t.Params {
		if text, err := wikicode.GetPlainText(); err == nil {
			if len(out) > 0 && len(text) > 0 {
				out += ", " + text
			}
		}
	}
	return out, nil
}

func (t *Template) handleHList() (string, error) {
	return t.handleList()
}

func (t *Template) handlePlainList() (string, error) {
	return t.handleList()
}

func (t *Template) handleFlatList() (string, error) {
	return t.handleList()
}

func (t *Template) handleOrderedList() (string, error) {
	return t.handleList()
}

func (t *Template) handleUnbulletedList() (string, error) {
	return t.handleList()
}

func (t *Template) handlePageList() (string, error) {
	return t.handleList()
}

// ----- URLs -----
func (t *Template) handleURL() (string, error) {
	return t.getParamTextByKeyIndex(1)
}

// ----- Height -----

// m, cm
func (t *Template) handleHeight() (string, error) {
	for key, symbol := range units {
		if wikicode, ok := t.Params[key]; ok && wikicode != nil {
			height, err := wikicode.GetPlainText()
			if err == nil {
				return fmt.Sprintf("%s %s", height, symbol), nil
			}
		}
	}
	return "", fmt.Errorf("could not handle height template: %s", t.Name)
}

func (t *Template) handleNowrap() (string, error) {
	return t.getParamTextByKeyIndex(1)
}

// -- utils
func (t *Template) getParamTextByKeyIndex(i int) (string, error) {
	var out string
	var err error
	if wikicode, ok := t.Params[fmt.Sprintf("%d", i)]; ok && wikicode != nil {
		out, err = wikicode.GetPlainText()
	}
	return out, err
}
