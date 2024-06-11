package core

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Path struct {
	Val string
}

type Corpus struct {
	Docs   []Doc
	Termos map[string]uint64
}

type corpus2save struct {
	DocsNames []string
	Termos    []string
}

func (c Corpus) Save() {

}

func CorpusFromFile(string) {

}

func (c Corpus) PrintCorpus() {
	fmt.Println("total termos ", len(c.Termos))
	fmt.Println("total Docs ", len(c.Docs))
}

func CorpusDirFactory(base Path, extension string) Corpus {
	all_paths := base.GetFiles(extension)
	all_docs := make([]Doc, 0, len(all_paths))
	termos := make(map[string]uint64)
	for _, v := range all_paths {
		d := FactoryDoc(v)
		for key, val := range d.Vocab {
			termos[key] += uint64(val)
		}
		all_docs = append(all_docs, d)
	}

	return Corpus{Docs: all_docs, Termos: termos}
}

func (p Path) GetFiles(extension string) []Path {

	return getFiles(p.Val, extension)
}

type kvPair struct {
	key   string
	value uint64
}

func (c Corpus) Show20() {
	var sortedTerms []kvPair

	for key, value := range c.Termos {
		sortedTerms = append(sortedTerms, kvPair{key: key, value: value})
	}
	sort.Slice(sortedTerms, func(i, j int) bool {
		return sortedTerms[i].value > sortedTerms[j].value
	})
	fileName := "top_terms.csv"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"Term", "Value"})
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}
	for _, term := range sortedTerms {
		err = writer.Write([]string{term.key, fmt.Sprintf("%d", term.value)})
		if err != nil {
			fmt.Println("Error writing data:", err)
			return
		}
	}

	fmt.Println("Salvo CSV com ", len(sortedTerms), " Termos.")
}

func getFiles(root string, extension string) []Path {
	paths := make([]Path, 0)
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)

	}
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getFiles(root+file.Name()+"/", extension)...)
		} else if strings.HasSuffix(file.Name(), extension) {
			paths = append(paths, Path{Val: root + file.Name()})
		}
	}
	return paths
}
