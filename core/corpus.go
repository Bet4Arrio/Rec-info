package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Path struct {
	Val string
}

type Corpus struct {
	Docs   []Doc
	Termos map[string]bool
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
	termos := make(map[string]bool)
	for _, v := range all_paths {
		d := FactoryDoc(v)
		for key := range d.Vocab {
			termos[key] = true
		}
		all_docs = append(all_docs, d)
	}

	return Corpus{Docs: all_docs, Termos: termos}
}

func (p Path) GetFiles(extension string) []Path {

	return getFiles(p.Val, extension)
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
