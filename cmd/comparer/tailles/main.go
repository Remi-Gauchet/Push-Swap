package main

import (
	"fmt"
	"math/rand"

	"github.com/Remi-Gauchet/Push-Swap/piles"
)

// comparerTailles teste plusieurs tailles de bloc pour chaque taille de pile
func comparerTailles() {
	tailles := []int{12, 20, 50, 100, 200, 500}
	blocs := []int{5, 8, 10, 12, 15, 20, 30}
	essais := 20

	for _, n := range tailles {
		fmt.Printf("n=%d : ", n)
		for _, b := range blocs {
			total := 0
			for e := 0; e < essais; e++ {
				p := piles.Piles{A: rand.Perm(n)}
				p.TrierParBlocs(b)
				total += len(p.Operations)
			}
			fmt.Printf("bloc %d : %d   ", b, total/essais)
		}
		fmt.Println()
	}
}

func main() {
	comparerTailles()
}
