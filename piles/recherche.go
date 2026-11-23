package piles

import (
	"fmt"

	"github.com/Remi-Gauchet/Push-Swap/outils"
)

func (p *Piles) copie() Piles {
	return Piles{
		A:          outils.Copier(p.A),
		B:          outils.Copier(p.B),
		Operations: outils.Copier(p.Operations),
	}
}

func (p *Piles) cle() string {
	return fmt.Sprint(p.A, p.B)
}

// Recherche renvoie la suite d'instructions la plus courte pour trier
func (p *Piles) Recherche() []string {
	instructions := []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

	depart := p.copie()
	file := []Piles{depart}
	vus := map[string]bool{depart.cle(): true} // le carnet des situations déjà rencontrées

	for len(file) > 0 {
		courant := file[0] // on prend la première de la file
		file = file[1:]    // et on l'enlève de la file

		// si courant est trié, on renvoie ses instructions
		if courant.AEstTriee() && len(courant.B) == 0 {
			return courant.Operations
		}

		for _, instruction := range instructions {
			suivant := courant.copie()
			suivant.Application(instruction)

			// on l'y note et on l'ajoute à la file
			cle := suivant.cle() // l'étiquette de cette situation
			if !vus[cle] {       // si elle n'est PAS dans le carnet
				file = append(file, suivant) //on l'ajoute à la fin de la file
				vus[cle] = true
			}
		}
	}
	return nil
}
