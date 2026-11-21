package outils

import (
	"math"
)

// TailleDeBloc renvoie une taille de bloc optimale pour trier une pile de n éléments (n >= 12)
// c'est à dire environ 1,4 × √n, valeur trouvée par mesure avec cmd/comparer
func TailleDeBloc(n int) int {
	var coefficient float64 = 1.4
	racine := math.Sqrt(float64(n))                 // racine carrée de n
	taille := int(math.Round(coefficient * racine)) // on arrondit la racine, puis on la convertit en int, puis on la multiplie par le coefficient
	return taille
}
