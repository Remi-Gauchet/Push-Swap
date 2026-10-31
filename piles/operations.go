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

//pile B : envoie le premier en dernier
func (p *Piles) rb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

// réalise ra et rb en même temps
func (p *Piles) rr() {
	if len(p.A) < 2 && len(p.B) > 1 {
		p.rb()
		return
	}
	if len(p.A) > 1 && len(p.B) < 2 {
		p.ra()
		return
	}
	if len(p.A) < 2 && len(p.B) < 2 {
		return
	}
	p.A = append(p.A[1:], p.A[0])
	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rr")
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
	if len(p.B) < 2 || len(p.A) < 2 {
		return
	}
	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "ss")
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
	if len(p.A) < 2 || len(p.B) < 2 {
		return
	}
	p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	p.Operations = append(p.Operations, "rrr")
}
