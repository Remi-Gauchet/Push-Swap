package piles

import "testing"

func TestRecherche(t *testing.T) {
	p := Piles{A: []int{7, 6, 5, 4, 3, 2, 1, 0}}
	resultat := p.Recherche()
	t.Log(resultat)
	if len(resultat) != 16 {
		t.Errorf("attendu 16 instructions, obtenu %d : %v", len(resultat), resultat)
	}
}
