package tree

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetPlainText(t *testing.T) {
	samples := map[*Template]string{
		// birth date and age
		{
			Name: "birth date and age",
			Params: map[string]*Wikicode{
				"1":  {list: []Elem{NewText("1987")}},
				"2":  {list: []Elem{NewText("6")}},
				"3":  {list: []Elem{NewText("24")}},
				"df": {list: []Elem{NewText("y")}}},
		}: "24-6-1987",
		{
			Name: "birth-date and age",
			Params: map[string]*Wikicode{
				"1": {list: []Elem{NewText("1987")}}},
		}: "1987",

		// height
		{
			Name: "height",
			Params: map[string]*Wikicode{
				"m": {list: []Elem{NewText("1,85")}}},
		}: "1,85 m",
		{
			Name: "height",
			Params: map[string]*Wikicode{
				"cm":   {list: []Elem{NewText("177")}},
				"abbr": {list: []Elem{NewText("no")}}},
		}: "177 cm",

		// URL
		{
			Name: "URL",
			Params: map[string]*Wikicode{
				"1": {list: []Elem{NewText("taylorswift.com")}}},
		}: "taylorswift.com",

		// nowrap
		{
			Name: "nowrap",
			Params: map[string]*Wikicode{
				"1": {list: []Elem{NewText("Nhà sản xuất thu âm")}}},
		}: "Nhà sản xuất thu âm",

		// plainlist
		{
			Name: "plainlist",
			Params: map[string]*Wikicode{
				"1": {list: []Elem{NewText("Austin Swift")}},
				"2": {list: []Elem{NewText("Marjorie Finlay")}}},
		}: "Austin Swift, Marjorie Finlay",

		// hlist
		{
			Name: "hlist",
			Params: map[string]*Wikicode{
				"1": {list: []Elem{NewText("Austin Swift")}},
				"2": {list: []Elem{NewText("Marjorie Finlay")}}},
		}: "Austin Swift, Marjorie Finlay",

		// flatlist
		{
			Name: "flatlist",
			Params: map[string]*Wikicode{
				"1": {list: []Elem{NewText("Austin Swift")}},
				"2": {list: []Elem{NewText("Marjorie Finlay")}}},
		}: "Austin Swift, Marjorie Finlay",
	}

	for temp, expected := range samples {
		template0 := NewTemplate()
		template0.Name = temp.Name
		template0.Params = temp.Params
		out, err := template0.GetPlainText()
		assert.Equal(t, err, nil)
		assert.Equal(t, out, expected)
	}
}
