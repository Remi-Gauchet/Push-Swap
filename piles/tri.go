package piles

import (
	"github.com/Remi-Gauchet/Push-Swap/outils"
)

// trier appelle la fonction tri la plus adaptée à la taille de la pile
func (p *Piles) Trier() {
	if p.AEstTriee() {
		return
	}
	if len(p.A) == 2 {
		p.sa()
	} else if len(p.A) == 3 {
		p.trierTrois()
	} else if len(p.A) <= 8 { // on s'arrête à 8 sinon ça dure trop longtemps
		instructions := p.Recherche()
		for _, instruction := range instructions {
			p.Application(instruction)
		}
	} else if len(p.A) < 12 {
		p.TrierOptimise()
	} else {
		p.TrierParBlocs(outils.TailleDeBloc(len(p.A)))
	}
}

// renvoie True si une pile est tirée par ordre croissant
func (p *Piles) AEstTriee() bool {
	for i := 0; i < len(p.A)-1; i++ {
		if p.A[i] > p.A[i+1] {
			return false
		}
	}
	return true
}

// Méthode pour trier une pile de trois int
func (p *Piles) trierTrois() {
	if p.AEstTriee() == true {
		return
	} else {
		if p.A[0] < p.A[1] { // 1er < 2ème
			if p.A[1] > p.A[2] { // 1er < 2ème et 2ème > 3ème : 132 ou 231
				if p.A[0] < p.A[2] {
					// 132
					p.sa()
					p.ra()
					return
				}
				//231
				p.rra()
				return
			}
		}
		// 1er > 2ème
		if p.A[1] < p.A[2] { // 1er > 2ème et 2eme < 3eme : 312 ou 213
			if p.A[0] < p.A[2] {
				// 213
				p.sa()
				return
			}
			// 312
			p.ra()
			return
		}

		if p.A[0] > p.A[2] { // 1er > 3ème
			// alors ça commence par 3
			if p.A[1] > p.A[2] { // 2ème > 3ème
				// 321
				p.sa()
				p.rra()
				return
			}
		}
	}
}

// trierCinq trie une pile de 5
func (p *Piles) trierCinq() {
	for len(p.A) > 3 {
		p.amenerAuSommetA(p.positionDuMinimum())
		p.pb()
	}
	p.trierTrois()
	for len(p.B) > 0 {
		p.pa()
	}
}

// trierOptimise trie par sélection
func (p *Piles) TrierOptimise() {
	for len(p.A) != 0 {
		p.amenerAuSommetA(p.positionDuMinimum())
		p.pb()
	}
	for len(p.B) != 0 {
		p.pa()
	}
}

// amenerAuSommetA fait remonter l'élément de la pile "A" à l'index "position",
// par le chemin le plus court
func (p *Piles) amenerAuSommetA(position int) {
	if position <= len(p.A)/2 {
		for i := 0; i < position; i++ {
			p.ra()
		}
	} else {
		for i := 0; i < len(p.A)-position; i++ {
			p.rra()
		}
	}
}

// amenerAuSommetB fait remonter l'élément de la pile "B" à l'index "position",
// par le chemin le plus court
func (p *Piles) amenerAuSommetB(position int) {
	if position <= len(p.B)/2 {
		for i := 0; i < position; i++ {
			p.rb()
		}
	} else {
		for i := 0; i < len(p.B)-position; i++ {
			p.rrb()
		}
	}
}

// remonterLesPlusGrands ramène sur la pile "a" les éléments de la pile "b"
// en commençant par les plus grands jusqu'aux au plus petits
func (p *Piles) remonterLesPlusGrands() {
	for len(p.B) > 0 {
		p.amenerAuSommetB(p.positionDuMaximumB())
		p.pa()
	}
}

// trierParBlocs trie a en deux phases : pousser par blocs, puis remonter.
func (p *Piles) TrierParBlocs(tailleBloc int) {
	p.pousserParBlocs(tailleBloc)
	p.remonterLesPlusGrands()
}
