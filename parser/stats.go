package parser

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type infobox map[string]string

var (
	records map[string][]*infobox = make(map[string][]*infobox)
	outDir  string                = "raw_data/wiki-infobox-stats"
	mu      sync.Mutex
)

func WriteLog() error {
	if err := os.RemoveAll(outDir); err != nil {
		logrus.Error(err)
		return err
	}
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		logrus.Error(err)
		return err
	}
	for templateName, infoboxes := range records {
		templateName = strings.ReplaceAll(templateName, " ", "_")
		writeAll2File(infoboxes, path.Join(outDir, strings.ReplaceAll(templateName, " ", "_")+"-all.o"))
		writeNames2File(infoboxes, path.Join(outDir, templateName+"-names.o"))
	}
	return nil
}

func writeAll2File(infoboxes []*infobox, filePath string) error {
	var err error
	var file *os.File
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for _, infobox := range infoboxes {
		name := (*infobox)["name"]
		if name == "" {
			name = (*infobox)["tên"]
		}
		writer.WriteString(fmt.Sprintf("%s\n", name))
		for key, val := range *infobox {
			writer.WriteString(fmt.Sprintf("\t%s: %s\n", key, val))
		}
		writer.WriteString("\n")
	}
	return nil
}

func writeNames2File(infoboxes []*infobox, filePath string) error {
	var err error
	var file *os.File
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for _, infobox := range infoboxes {
		name := (*infobox)["name"]
		if name == "" {
			name = (*infobox)["tên"]
		}
		if len(name) == 0 {
			continue
		}
		writer.WriteString(fmt.Sprintf("%s\n", name))
	}
	return nil
}
