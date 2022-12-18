package templates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"wikitext-parser/utils"
)

var (
	mapTemplateName2Id            map[string]int
	mapTemplateId2Fields          map[int][]*field
	mapTemplateId2SubjectCategory map[int]string
)

type field struct {
	Id            int    `json:"id"`
	En            string `json:"en"`
	Vi            string `json:"vi"`
	DisplayedText string `json:"displayed-text"`
	Enabled       bool   `json:"enabled"`
}

type template struct {
	EnLabel         string   `json:"en-label"`
	ViLabel         string   `json:"vi-label"`
	SubjectCategory string   `json:"subject-category"`
	Fields          []*field `json:"fields"`
}

type templates []template

// read templates.json and build maps
func init() {
	mapTemplateName2Id = make(map[string]int)
	mapTemplateId2Fields = make(map[int][]*field)
	mapTemplateId2SubjectCategory = make(map[int]string)

	// TODO: modify the path to templates.json
	jsonFile, err := os.Open("/home/annv/go/src/wikitext-parser/templates/templates.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var temps templates
	json.Unmarshal(byteValue, &temps)

	for i, t := range temps {
		mapTemplateName2Id[utils.PreprocessTemplateName(t.EnLabel)] = i
		mapTemplateName2Id[utils.PreprocessTemplateName(t.ViLabel)] = i
		mapTemplateId2Fields[i] = t.Fields
		mapTemplateId2SubjectCategory[i] = t.SubjectCategory
	}
	log.Print("Loading templates.json done.")
}

func GetFieldsFromTemplate(templateName string) []*field {
	if fieldId, ok := mapTemplateName2Id[templateName]; ok {
		return mapTemplateId2Fields[fieldId]
	}
	return nil
}

func GetSubjectCategoryFromTemplate(templateName string) string {
	if tempId, ok := mapTemplateName2Id[templateName]; ok {
		return mapTemplateId2SubjectCategory[tempId]
	}
	return ""
}
