package outils

import (
	"fmt"
	"strconv"
	"strings"
)

// LireNombres récupère les arguments et les convertit en int
func LireNombres(args []string) ([]int, error) {
	var nombres []int
	for _, texte := range args {
		morceauxTexte := strings.Fields(texte)
		for j := range morceauxTexte {
			valeur, err := strconv.Atoi(morceauxTexte[j])
			if err != nil {
				return nil, fmt.Errorf("%q n'est pas un nombre entier", morceauxTexte[j])
			}
			nombres = append(nombres, valeur)
		}
	}
	return nombres, nil
}
