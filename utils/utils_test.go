package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPreprocessTemplateName(t *testing.T) {
	samples := [][]string{
		{"hộp thông tin quốc gia", "quốc gia"},
		{"infobox country", "country"},
		{"thông tin diễn viên", "diễn viên"},
	}
	for _, sample := range samples {
		assert.Equal(t, len(sample), 2)
		templateName := sample[0]
		expected := sample[1]
		assert.Equal(t, PreprocessTemplateName(templateName), expected)
	}
}
