package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/bet4arrio/Inforecs/core"
	"github.com/bet4arrio/Inforecs/vetorial"
)

func main() {

	fmt.Println("Hello word test")
	a := core.Path{Val: "/home/emanu/media/IF/REC-info/gosys/data/"}
	start := time.Now()
	startbase := time.Now()
	corpos := core.CorpusDirFactory(a, ".txt")
	fmt.Printf("%s took %v\n", "Corpus", time.Since(start))
	corpos.PrintCorpus()
	start = time.Now()
	sis := vetorial.VetorialFactory(corpos)
	fmt.Printf("%s took %v\n", "Vetorial", time.Since(start))
	fmt.Printf("%s took %v\n", "corpus+Vetorial", time.Since(startbase))
	// start = time.Now()
	// sis.Save()
	// fmt.Printf("%s took %v\n", "save", time.Since(start))
	// start = time.Now()

	// sis2 := vetorial.LoadFromJson("Vetorial.json")
	// fmt.Printf("%s took %v\n", "Load", time.Since(start))
	var q string
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Pesquisa: ")
	q, err := in.ReadString('\n')
	if err != nil {
		panic("panic")
	}
	for q != "cu\n" {
		fmt.Println(q)
		start = time.Now()
		teste := sis.Performquery(q[:len(q)-1])
		// teste2 := sis2.Performquery(q[:len(q)-1])
		fmt.Println(len(teste), "Encontrado")
		fmt.Print("Pesquisa: ")
		fmt.Printf("%s took %v\n", "Vetorial Save", time.Since(start))
		q, err = in.ReadString('\n')
		if err != nil {
			panic("panic")
		}
	}

}
