package main

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
	} else {
		p.trierOptimise()
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

// amenerAuSommet fait remonter l'élément de a en position donnée,
// par le chemin le plus court.
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

func (p *Piles) trierOptimise() {
	p.trierCinq()
}
