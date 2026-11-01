package outils

// TailleDeBloc choisit la taille des blocs selon le nombre d'éléments
func TailleDeBloc(n int) int {
	if n <= 100 {
		return 15
	}
	return 30
}
