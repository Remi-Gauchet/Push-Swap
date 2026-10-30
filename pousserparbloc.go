package main

// pousserParBlocs vide a dans b, en ne poussant que les rangs
// du bloc en cours. Les plus petits du bloc partent au fond de b.
func (p *Piles) pousserParBlocs(tailleBloc int) {
	pousses := 0

	for len(p.a) > 0 {
		// à compléter avec la règle de la partie 1
		if p.a[0] <= pousses {
			p.pb()
			p.rb()
			pousses++
		} else if p.a[0] <= (pousses + tailleBloc) {
			p.pb()
			pousses++
		} else {
			p.ra()
		}
	}
}
