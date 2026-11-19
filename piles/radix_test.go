package piles

import (
	"reflect"
	"testing"
)

func TestTrierRadix(t *testing.T) {
	cas := []struct {
		nom     string
		depart  []int
		attendu []int
	}{
		{
			nom:     "désordre simple",
			depart:  []int{3, 0, 2, 1},
			attendu: []int{0, 1, 2, 3},
		},
		{
			nom:     "ordre inverse",
			depart:  []int{4, 3, 2, 1, 0},
			attendu: []int{0, 1, 2, 3, 4},
		},
		{
			nom:     "déjà trié",
			depart:  []int{0, 1, 2, 3, 4},
			attendu: []int{0, 1, 2, 3, 4},
		},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			p := Piles{A: c.depart}

			p.TrierRadix()

			if !reflect.DeepEqual(p.A, c.attendu) {
				t.Errorf("obtenu %v, attendu %v", p.A, c.attendu)
			}

			if len(p.B) != 0 {
				t.Errorf("B devrait être vide, obtenu %v", p.B)
			}
		})
	}
}
