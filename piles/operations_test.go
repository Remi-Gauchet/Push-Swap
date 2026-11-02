package piles

import (
	"reflect"
	"testing"
)

func TestRa(t *testing.T) {
	p := Piles{A: []int{1, 2, 3}, B: []int{4, 5}}
	p.ra()
	valeurAttendue := Piles{A: []int{2, 3, 1}, B: []int{4, 5}, Operations: []string{"ra"}}
	if !reflect.DeepEqual(p, valeurAttendue) {
		t.Errorf("message : obtenu %v, attendu %v", p, valeurAttendue)
	}
}

func TestRr(t *testing.T) {
	cas := []struct {
		nom     string
		depart  Piles
		attendu Piles
	}{
		{
			nom:     "les deux piles tournent",
			depart:  Piles{A: []int{1, 2, 3}, B: []int{4, 5, 6}, Operations: []string{"pb"}},
			attendu: Piles{A: []int{2, 3, 1}, B: []int{5, 6, 4}, Operations: []string{"pb", "rr"}},
		},
		{
			nom:     "rien ne bouge",
			depart:  Piles{A: []int{1}, B: []int{4}, Operations: []string{"pb"}},
			attendu: Piles{A: []int{1}, B: []int{4}, Operations: []string{"pb"}},
		},
		{
			nom:     "seul A tourne",
			depart:  Piles{A: []int{1, 2, 3}, B: []int{4}, Operations: []string{"pb"}},
			attendu: Piles{A: []int{2, 3, 1}, B: []int{4}, Operations: []string{"pb", "ra"}},
		},
		{
			nom:     "seul B tourne",
			depart:  Piles{A: []int{1}, B: []int{4, 5, 6}, Operations: []string{"pb"}},
			attendu: Piles{A: []int{1}, B: []int{5, 6, 4}, Operations: []string{"pb", "rb"}},
		},
	}
	for _, c := range cas {
		p := c.depart
		p.rr()
		if !reflect.DeepEqual(p, c.attendu) {
			t.Errorf("%s : obtenu %+v, attendu %+v", c.nom, p, c.attendu)
		}
	}
}
