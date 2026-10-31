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

// comparer lance les deux stratégies sur les mêmes piles au hasard
// et affiche la moyenne des mouvements pour chaque taille.
func comparer() {
	tailles := []int{10, 20, 50, 100, 500}
	essais := 20

	for _, n := range tailles {
		totalSimple := 0
		totalBlocs := 0

		for e := 0; e < essais; e++ {
			depart := rand.Perm(n)
			copie1 := copier(depart)
			// à compléter : lancer trierOptimise sur une copie,
			Tsimple := piles.Piles{A: copie1}
			Tsimple.TrierOptimise()
			// trierParBlocs sur une autre copie,
			copie2 := copier(depart)
			Tbloc := piles.Piles{A: copie2}
			Tbloc.TrierParBlocs(outils.TailleDeBloc(n))
			// et ajouter leurs nombres d'opérations aux totaux
			totalSimple += len(Tsimple.Operations)
			totalBlocs += len(Tbloc.Operations)

		}

		moyenneSimple := totalSimple / essais
		moyenneBlocs := totalBlocs / essais
		gain := 100 * (moyenneSimple - moyenneBlocs) / moyenneSimple

		fmt.Printf("n=%-4d  séance 2 : %-6d  blocs : %-6d  gain : %d %%\n", n, moyenneSimple, moyenneBlocs, gain)
	}
}

func main() {
	comparer()
}
