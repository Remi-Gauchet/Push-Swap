package piles

// TrierRadix trie la pile A avec un tri binaire appelé tri radix
func (p *Piles) TrierRadix() {
	if len(p.A) < 2 {
		return
	}

	// Trouver le nombre de bits nécessaire pour représenter le plus grand rang
	max := len(p.A) - 1
	nombreBits := 0

	for max > 0 {
		max >>= 1
		nombreBits++
	}

	// On traite les bits du moins significatif au plus significatif
	for bit := 0; bit < nombreBits; bit++ {
		taille := len(p.A)

		for i := 0; i < taille; i++ {
			// Si le bit vaut 0 : on push sur B
			if (p.A[0]>>bit)&1 == 0 {
				p.pb()
			} else {
				// Si le bit vaut 1 : on le bazarde à l'arrière
				p.ra()
			}
		}

		// On remet tous les éléments dans A
		for len(p.B) > 0 {
			p.pa()
		}
	}
}
