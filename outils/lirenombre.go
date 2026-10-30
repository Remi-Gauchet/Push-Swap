package outils

import (
	"fmt"
	"strconv"
)

// lireNombre récupère les arguments et les convertir en int
func LireNombres(args []string) ([]int, error) {
	var nombres []int
	for i := range args {
		valeur, err := strconv.Atoi(args[i])
		if err != nil {
			return nil, fmt.Errorf("%q n'est pas un nombre entier", args[i])
		}
		nombres = append(nombres, valeur)
	}
	return nombres, nil
}
