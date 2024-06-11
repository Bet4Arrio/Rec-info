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

	fmt.Println("Vetor Search")
	a := core.Path{Val: "./data/"}
	start := time.Now()
	startbase := time.Now()
	fmt.Println("mounting Corpus")
	corpos := core.CorpusDirFactory(a, ".txt")
	fmt.Printf("%s levou %v\n", "Corpus", time.Since(start))
	corpos.PrintCorpus()
	corpos.Show20()
	start = time.Now()
	fmt.Println("mounting Vetorial Model")
	fmt.Println("mounting Vetorial Model")
	sis := vetorial.VetorialFactory(corpos)
	fmt.Printf("%s levou %v\n", "Vetorial", time.Since(start))
	fmt.Printf("%s levou %v\n", "corpus+Vetorial", time.Since(startbase))
	var q string
	in := bufio.NewReader(os.Stdin)
	fmt.Println("===========================================================")
	fmt.Println("VETORIAL SEATCH UwU")
	fmt.Println("===========================================================")
	fmt.Print("Pesquisa: ")
	q, err := in.ReadString('\n')
	if err != nil {
		panic("panic")
	}
	for q != ":q\n" {
		fmt.Println(q)
		start = time.Now()
		teste := sis.Performquery(q[:len(q)-1])
		fmt.Println(len(teste), "Encontrados")
		fmt.Printf("%s levou %v\n", "Pesquisa", time.Since(start))
		vetorial.PageQuery(teste)
		fmt.Print("Nova Pesquisa(:q para sair): ")
		q, err = in.ReadString('\n')
		if err != nil {
			panic("panic")
		}
	}
}
