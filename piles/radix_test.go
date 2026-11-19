package piles

import (
	"fmt"
	"math/rand"
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

func TestTrierRadixAleatoire(t *testing.T) {
	tailles := []int{5, 20, 100}

	rng := rand.New(rand.NewSource(42))

	for _, taille := range tailles {
		t.Run(fmt.Sprintf("%d éléments", taille), func(t *testing.T) {

			for i := 0; i < 100; i++ {

				// Générer une permutation aléatoire
				depart := rng.Perm(taille)

				p := Piles{A: depart}

				// Exécuter le radix
				p.TrierRadix()

				// Vérifier que A est triée
				for j := 0; j < len(p.A)-1; j++ {
					if p.A[j] > p.A[j+1] {
						t.Fatalf(
							"échec avec %v : A = %v",
							depart,
							p.A,
						)
					}
				}

				// Vérifier que B est vide
				if len(p.B) != 0 {
					t.Fatalf("B devrait être vide : %v", p.B)
				}
			}
		})
	}
}
