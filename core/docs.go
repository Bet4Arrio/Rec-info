package core

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Doc struct {
	Name  Path
	Vocab map[string]uint32
}

func FactoryDoc(fname Path) Doc {
	content, err := os.ReadFile(fname.Val)
	if err != nil {
		log.Fatal(err, fname)
	}
	d := Doc{Name: fname, Vocab: VocabSetCounter(CleanUp(string(content)))}
	return d
}

func CleanUp(q string) string {
	return strings.ToLower(q)
}

func VocabSetCounter(q string) map[string]uint32 {
	Vocab := make(map[string]uint32)
	for _, v := range GetTokens(q) {
		Vocab[v]++
	}
	return Vocab
}
func GetTokens(q string) []string {
	rgx := regexp.MustCompile(`\w+`)
	return rgx.FindAllString(q, -1)
}
