package core

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Doc struct {
	Name    Path
	Content string
	Vocab   map[string]bool
}

func FactoryDoc(fname Path) Doc {
	content, err := os.ReadFile(fname.Val)
	if err != nil {
		log.Fatal(err, fname)
	}
	d := Doc{Name: fname, Content: strings.ToLower(string(content)), Vocab: make(map[string]bool)}
	d.VocabSet()
	return d
}

func (d Doc) VocabSet() {
	for _, v := range d.GetTokens() {
		d.Vocab[v] = true
	}
}

func (d Doc) GetTokens() []string {
	return GetTokens(strings.ReplaceAll(d.Content, "\n", ""))
}
func GetTokens(q string) []string {
	rgx := regexp.MustCompile(`\w+`)
	return rgx.FindAllString(q, -1)
}
