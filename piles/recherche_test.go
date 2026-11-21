package piles

import "testing"

func TestRecherche(t *testing.T) {
	p := Piles{A: []int{1, 3, 0, 4, 2}}
	resultat := p.Recherche()
	t.Log(resultat)
	if len(resultat) != 6 {
		t.Errorf("attendu 6 instructions, obtenu %d : %v", len(resultat), resultat)
	}
}
