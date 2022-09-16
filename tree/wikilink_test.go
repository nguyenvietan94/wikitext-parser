package tree

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestWikilinkGetPlainText(t *testing.T) {
	samples := map[*Wikilink]string{
		{wikiPage: "Tập tin:Apple logo black.svg", displayedText: ""}:  "Apple logo black.svg",
		{wikiPage: ":Tập tin:Apple logo black.svg", displayedText: ""}: "Apple logo black.svg",
		{wikiPage: "File:Apple logo black.svg", displayedText: ""}:     "Apple logo black.svg",
		{wikiPage: ":File:Apple logo black.svg", displayedText: ""}:    "Apple logo black.svg",
		{wikiPage: "Media:Apple logo black.svg", displayedText: ""}:    "Apple logo black.svg",
	}
	for wikilink, expected := range samples {
		if text, err := wikilink.GetPlainText(); err == nil {
			assert.Equal(t, text, expected)
		}
	}
}
