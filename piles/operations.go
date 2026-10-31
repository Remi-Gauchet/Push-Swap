package piles

// Détail des opérations

// pile A : envoie le premier en dernier
func (p *Piles) ra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append(p.A[1:], p.A[0])
	p.Operations = append(p.Operations, "ra")
}

// pile B : envoie le premier en dernier
func (p *Piles) rb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

// réalise ra et rb en même temps
func (p *Piles) rr() {
	switch {
	case len(p.A) < 2:
		p.rb()
	case len(p.B) < 2:
		p.ra()
	default:
		p.ra()
		p.rb()
		p.Operations = append(p.Operations[:len(p.Operations)-2], "rr")
	}
}

// échange les deux premiers éléments de la pile A
func (p *Piles) sa() {
	if len(p.A) < 2 {
		return
	}
	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.Operations = append(p.Operations, "sa")
}

// échange les deux premiers éléments de la pile B
func (p *Piles) sb() {
	if len(p.B) < 2 {
		return
	}
	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "sb")
}

// réalise sa et sb en même temps
func (p *Piles) ss() {
	switch {
	case len(p.B) < 2:
		p.sa()
	case len(p.A) < 2:
		p.sb()
	default:
		p.sa()
		p.sb()
		p.Operations = append(p.Operations[:len(p.Operations)-2], "ss")
	}
}

// envoie le premier élément de la pile A sur la pile B
func (p *Piles) pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

// envoie le premier élément de la pile B sur la pile A
func (p *Piles) pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

// Pile A : envoie le dernier en premier
func (p *Piles) rra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	p.Operations = append(p.Operations, "rra")
}

// Pile B : envoie le dernier en premier
func (p *Piles) rrb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	p.Operations = append(p.Operations, "rrb")
}

// exécute rra et rrb en même temps
func (p *Piles) rrr() {
	switch {
	case len(p.A) < 2:
		p.rrb()
	case len(p.B) < 2:
		p.rra()
	default:
		p.rra()
		p.rrb()
		p.Operations = append(p.Operations[:len(p.Operations)-2], "rrr")
	}
}
