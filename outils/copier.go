package outils

// copier renvoie une copie indépendante de la tranche.
func Copier(nombres []int) []int {
	copie := make([]int, len(nombres))
	copy(copie, nombres)
	return copie
}
