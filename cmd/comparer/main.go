package main

import (
	"fmt"
	"math/rand"

	"github.com/Remi-Gauchet/Push-Swap/outils"
	"github.com/Remi-Gauchet/Push-Swap/piles"
)

// copier renvoie une copie indépendante de la tranche.
func copier(nombres []int) []int {
	copie := make([]int, len(nombres))
	copy(copie, nombres)
	return copie
}

// comparer lance nos 3 stratégies de tri sur les mêmes piles tirées au hasard
// TrierOptimise, TrierParBlocs et TrierRadix (détails dans le README)
// puis affiche la moyenne des mouvements pour chaque taille, afin de comparer leur efficacité en fonction de n
func comparer() {
	tailles := []int{10, 20, 50, 100, 500, 1000}
	essais := 20

	for _, n := range tailles {
		totalSimple := 0
		totalBlocs := 0
		totalRadix := 0

		for e := 0; e < essais; e++ {
			depart := rand.Perm(n)

			// Tri par sélection
			copie1 := copier(depart)
			Tsimple := piles.Piles{A: copie1}
			Tsimple.TrierOptimise()

			// tri par blocs
			copie2 := copier(depart)
			Tbloc := piles.Piles{A: copie2}
			Tbloc.TrierParBlocs(outils.TailleDeBloc(n))

			// tri radix
			copie3 := copier(depart)
			Tradix := piles.Piles{A: copie3}
			Tradix.TrierRadix()

			//addition
			totalSimple += len(Tsimple.Operations)
			totalBlocs += len(Tbloc.Operations)
			totalRadix += len(Tradix.Operations)
		}

		// calcul des moyennes
		moyenneSimple := totalSimple / essais
		moyenneBlocs := totalBlocs / essais
		moyenneRadix := totalRadix / essais

		gainBlocs := 0
		gainRadix := 0

		if moyenneSimple > 0 {
			gainBlocs = (moyenneSimple - moyenneBlocs) * 100 / moyenneSimple
			gainRadix = (moyenneSimple - moyenneRadix) * 100 / moyenneSimple
		}

		fmt.Printf(
			"n=%-4d  sélection : %-6d  blocs : %-6d (gain blocs : %d%%)  radix : %-6d (gain radix : %d%%)\n",
			n,
			moyenneSimple,
			moyenneBlocs,
			gainBlocs,
			moyenneRadix,
			gainRadix,
		)
	}
}

func main() {
	comparer()
}
