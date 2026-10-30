package main

import (
	"fmt"
	"math/rand"

	"github.com/Remi-Gauchet/Algorithmie/seance5/outils"
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
			simple := Piles{a: copie1}
			simple.trierOptimise()
			// trierParBlocs sur une autre copie,
			copie2 := copier(depart)
			bloc := Piles{a: copie2}
			bloc.trierParBlocs(outils.TailleDeBloc(n))
			// et ajouter leurs nombres d'opérations aux totaux
			totalSimple += len(simple.operations)
			totalBlocs += len(bloc.operations)

		}

		moyenneSimple := totalSimple / essais
		moyenneBlocs := totalBlocs / essais
		gain := 100 * (moyenneSimple - moyenneBlocs) / moyenneSimple

		fmt.Printf("n=%-4d  séance 2 : %-6d  blocs : %-6d  gain : %d %%\n", n, moyenneSimple, moyenneBlocs, gain)
	}
}
