package outils

// Copier renvoie une copie indépendante de la slice
// prend en compte les []int et les []string
func Copier[T any](nombres []T) []T {
	copie := make([]T, len(nombres))
	copy(copie, nombres)
	return copie
}
