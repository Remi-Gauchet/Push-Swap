package outils

import "fmt"

// renvoie une erreur s'il y a des doublons dans la liste d'int
func VerifierDoublons(nombres []int) error {
	vus := map[int]bool{}
	for _, nombre := range nombres {
		if vus[nombre] {
			// déjà vu : c'est un doublon
			return fmt.Errorf("le nombre %d apparaît plusieurs fois", nombre)
		}
		vus[nombre] = true
	}
	return nil
}
