package piles

// pousserParBlocs vide a dans b, en ne poussant que les rangs
// du bloc en cours. Les plus petits du bloc partent au fond de b.
func (p *Piles) pousserParBlocs(tailleBloc int) {
	pousses := 0
	for len(p.A) > 0 {
		if p.A[0] <= pousses {
			p.pb()
			pousses++
			if len(p.A) > 0 && p.A[0] > (pousses+tailleBloc) {
				p.rr()
			} else {
				p.rb()
			}
		} else if p.A[0] <= (pousses + tailleBloc) {
			pousses++
			p.pb()
		} else {
			p.ra()
		}
	}
}
