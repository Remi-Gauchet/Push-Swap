package main

import (
	"fmt"
	"math/rand"

	"github.com/Remi-Gauchet/Push-Swap/outils"
	"github.com/Remi-Gauchet/Push-Swap/piles"
)

// comparerRecherche compare une méthode de tri dédiée aux petites piles
// avec la recherche exhaustive, sur des piles de n éléments tirées au hasard
func comparerRecherche(nom string, n int, trier func(*piles.Piles)) {
	essais := 1000
	totalMethode := 0
	totalRecherche := 0

	for e := 0; e < essais; e++ {
		depart := rand.Perm(n)

		p := piles.Piles{A: outils.Copier(depart)}
		trier(&p) // on applique la méthode reçue en paramètre

		pr := piles.Piles{A: outils.Copier(depart)}
		instructions := pr.Recherche()

		totalMethode += len(p.Operations)
		totalRecherche += len(instructions)
	}

	moyenneMethode := float64(totalMethode) / float64(essais)
	moyenneRecherche := float64(totalRecherche) / float64(essais)

	fmt.Printf("Nombre moyen d'opérations pour 1000 essais avec n=%d  %s : %.2f  recherche : %.2f\n", n, nom, moyenneMethode, moyenneRecherche)
}
