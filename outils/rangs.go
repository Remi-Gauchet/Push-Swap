package outils

func Rangs(nombres []int) []int {
	var resultat []int
	for i := range nombres {
		var compteur int
		for j := range nombres {
			if nombres[i] > nombres[j] {
				compteur++
			}
		}
		resultat = append(resultat, compteur)
	}
	return resultat
}
