package main

import (
	"github.com/Remi-Gauchet/Algorithmie/seance5/outils"
)

// trier appelle la fonction tri la plus adaptée à la taille de la pile
func (p *Piles) trier() {
	if p.aEstTriee() {
		return
	}

	if len(p.a) == 2 {
		p.sa()
	} else if len(p.a) == 3 {
		p.trierTrois()
	} else if len(p.a) <= 5 {
		p.trierCinq()
	} else if len(p.a) < 20 {
		p.trierOptimise()
	} else {
		p.trierParBlocs(outils.TailleDeBloc(len(p.a)))
	}
}

// renvoie True si une pile est tirée par ordre croissant
func (p *Piles) aEstTriee() bool {
	for i := 0; i < len(p.a)-1; i++ {
		if p.a[i] > p.a[i+1] {
			return false
		}
	}
	return true
}

// Méthode pour trier une pile de trois int
func (p *Piles) trierTrois() {
	if p.aEstTriee() == true {
		return
	} else {
		if p.a[0] < p.a[1] { // 1er < 2ème
			if p.a[1] > p.a[2] { // 1er < 2ème et 2ème > 3ème : 132 ou 231
				if p.a[0] < p.a[2] {
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
		if p.a[1] < p.a[2] { // 1er > 2ème et 2eme < 3eme : 312 ou 213
			if p.a[0] < p.a[2] {
				// 213
				p.sa()
				return
			}
			// 312
			p.ra()
			return
		}

		if p.a[0] > p.a[2] { // 1er > 3ème
			// alors ça commence par 3
			if p.a[1] > p.a[2] { // 2ème > 3ème
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
	for len(p.a) > 3 {
		p.amenerAuSommet(p.positionDuMinimum())
		p.pb()
	}
	p.trierTrois()
	for len(p.b) > 0 {
		p.pa()
	}
}

// trierOptimise trie par sélection
func (p *Piles) trierOptimise() {
	for len(p.a) != 0 {
		p.amenerAuSommet(p.positionDuMinimum())
		p.pb()
	}
	for len(p.b) != 0 {
		p.pa()
	}
}

// amenerAuSommet fait remonter l'élément de la pile "a" à l'index "position",
// par le chemin le plus court
func (p *Piles) amenerAuSommet(position int) {
	if position < len(p.a)/2 {
		for i := 0; i < position; i++ {
			p.ra()
		}
	} else {
		for i := 0; i < len(p.a)-position; i++ {
			p.rra()
		}
	}
}

// amenerAuSommetB fair remonter l'élément de la pile "b" à l'index "position",
// par le chemin le plus court
func (p *Piles) amenerAuSommetB(position int) {
	if position < len(p.b)/2 {
		for i := 0; i < position; i++ {
			p.rb()
		}
	} else {
		for i := 0; i < len(p.b)-position; i++ {
			p.rrb()
		}
	}
}

// remonterLesPlusGrands ramène sur la pile "a" les éléments de la pile "b"
// en commençant par les plus grands jusqu'aux au plus petits
func (p *Piles) remonterLesPlusGrands() {
	for len(p.b) > 0 {
		p.amenerAuSommetB(p.positionDuMaximumB())
		p.pa()
	}
}

// trierParBlocs trie a en deux phases : pousser par blocs, puis remonter.
func (p *Piles) trierParBlocs(tailleBloc int) {
	p.pousserParBlocs(tailleBloc)
	p.remonterLesPlusGrands()
}
