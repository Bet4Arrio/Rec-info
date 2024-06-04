package vetorial

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/bet4arrio/Inforecs/core"
)

type SistemaVetorial struct {
	BoW        [][]uint32        `json:"bow"`    // DOC Incidencia TERMO
	Tf         [][]float32       `json:"tf"`     // DOC Incidencia TERMO
	TfIDF      [][]float32       `json:"tf-idf"` // DOC Incidencia TERMO
	Docs_names []string          `json:"docs"`
	Termos     map[string]uint64 `json:"termos"`
	// TermosMap
}

type QueriedDocs struct {
	Name string
	Rank float64
}

func LoadFromJson(name string) SistemaVetorial {
	var loadedSys SistemaVetorial
	jsonData, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonData, &loadedSys)
	if err != nil {
		panic(err)
	}

	return loadedSys

}

func VetorialFactory(c core.Corpus) SistemaVetorial {
	var Ndocs, Ntermos int
	Ndocs = len(c.Docs)
	Ntermos = len(c.Termos)
	Matriz := make([][]uint32, Ndocs)
	for i := range Matriz {
		Matriz[i] = make([]uint32, Ntermos)
	}

	docsName := make([]string, 0, Ndocs)
	termos := make(map[string]uint64, Ntermos)
	for _, v := range c.Docs {
		docsName = append(docsName, v.Name.Val)
	}
	conter := 0
	for k := range c.Termos {
		for i, d := range c.Docs {
			qnt, ok := d.Vocab[k]
			if ok {
				Matriz[i][conter] += qnt
			}
		}
		// fmt.Println(Ndocs, Ntermos, Matriz[1][:])
		termos[k] = uint64(conter)
		conter++
	}

	return SistemaVetorial{
		BoW:        Matriz,
		Tf:         calcTf(Matriz),
		TfIDF:      calcTfIDF(Matriz),
		Termos:     termos,
		Docs_names: docsName,
	}
}

func calcTf(bow [][]uint32) [][]float32 {
	termFreqs := make([][]float32, len(bow))
	for i, doc := range bow {
		termFreqs[i] = make([]float32, len(doc))
		totalTerms := float32(0)
		for _, termCount := range doc {
			totalTerms += float32(termCount)
		}
		for j, count := range doc {
			termFreqs[i][j] = float32(count) / totalTerms
		}
	}

	return termFreqs
}

func calcIDF(bow [][]uint32) [][]float32 {
	Matriz := make([][]float32, len(bow))
	for i := range Matriz {
		Matriz[i] = make([]float32, len(bow[i]))
		for j := range Matriz[i] {
			if bow[i][j] > 0 {
				Matriz[i][j] = float32(math.Log2(float64(bow[i][j]))) + 1

			}
		}
	}

	return Matriz
}

func calcTfIDF(bow [][]uint32) [][]float32 {
	// Check if documents is zero to avoid division by zero
	documents := len(bow)
	if documents == 0 {
		panic("Number of documents cannot be zero")
	}

	// Initialize TF-IDF matrix with zeros
	tfidfMatrix := make([][]float32, len(bow))
	for i := range tfidfMatrix {
		tfidfMatrix[i] = make([]float32, len(bow[i]))
	}

	// Calculate term frequency (TF) for each document and term
	termFreqs := make([][]float32, len(bow))
	for i, doc := range bow {
		termFreqs[i] = make([]float32, len(doc))
		totalTerms := float32(0)
		for _, termCount := range doc {
			totalTerms += float32(termCount)
		}
		for j, count := range doc {
			termFreqs[i][j] = float32(count) / totalTerms
		}
	}

	// Calculate inverse document frequency (IDF) for each term
	idf := make([]float32, len(bow[0]))
	for term := range bow[0] {
		documentFrequency := 0
		for _, doc := range bow {
			if doc[term] > 0 {
				documentFrequency++
			}
		}
		idf[term] = float32(math.Log(float64(documents) / float64(documentFrequency)))
	}

	// Calculate TF-IDF by multiplying TF and IDF for each term in each document
	for i, doc := range termFreqs {
		for j, tf := range doc {
			tfidfMatrix[i][j] = tf * idf[j]
		}
	}

	return tfidfMatrix
}

func (s SistemaVetorial) Performquery(q string) []QueriedDocs {
	vector := s.bowOfQ(q)
	resp := make([]QueriedDocs, 0, len(s.Docs_names))
	for i := range s.BoW {
		rank := calcCosProx(vector, s.BoW[i])
		if rank > 0 {
			resp = append(resp, QueriedDocs{Name: s.Docs_names[i], Rank: rank})

		}
	}

	return resp

}

func (s SistemaVetorial) bowOfQ(query string) []uint32 {

	bow := make([]uint32, len(s.Termos))
	termos := core.GetTokens(query)
	for i := range termos {
		t, ok := s.Termos[termos[i]]
		if ok {
			bow[t]++
		}
	}
	return bow
}

// func (s SistemaVetorial) tfIdfOfQ(query string) []float32 {
// 	base := make([]float32, len(s.Termos))
// 	termos := core.GetTokens(query)
// 	for i := range termos {
// 		t, ok := s.Termos[termos[i]]
// 		if ok {
// 			base[t]++
// 		}
// 	}

//		return base
//	}
func calcCosProx(q, v interface{}) float64 {
	// Type assertion to convert interface{} to concrete types

	dotProduct := 0.0
	normVector1 := 0.0
	normVector2 := 0.0
	switch q := q.(type) {
	case []float32:
		switch v := v.(type) {
		case []float32:
			for i := range q {
				dotProduct += float64(q[i] * v[i])
				normVector1 += float64(q[i] * q[i])
				normVector2 += float64(v[i] * v[i])
			}
		case []uint32:
			for i := range q {
				dotProduct += float64(q[i]) * float64(v[i])
				normVector1 += float64(q[i] * q[i])
				normVector2 += float64(v[i] * v[i])
			}
		}
	case []uint32:
		switch v := v.(type) {
		case []float32:
			for i := range q {
				dotProduct += float64(q[i]) * float64(v[i])
				normVector1 += float64(q[i] * q[i])
				normVector2 += float64(v[i] * v[i])
			}
		case []uint32:
			for i := range q {
				dotProduct += float64(q[i]) * float64(v[i])
				normVector1 += float64(q[i] * q[i])
				normVector2 += float64(v[i] * v[i])
			}
		}
	default:
		panic("invlaid input")
	}

	similarity := dotProduct / (math.Sqrt(normVector1) * math.Sqrt(normVector2))
	return similarity
}

func (s SistemaVetorial) Save() {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("Vetorial.json", data, 0777) // Adjust permissions as needed
	if err != nil {
		panic(err)
	}
	fmt.Println("Salvando")
}
