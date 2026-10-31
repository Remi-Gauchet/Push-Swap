package piles

// définition des opérations à exécuter sur les piles (à part du main.go pour alléger)
func (p *Piles) ra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append(p.A[1:], p.A[0])
	p.Operations = append(p.Operations, "ra")
}

func (p *Piles) rb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

func (p *Piles) sa() {
	if len(p.A) < 2 {
		return
	}
	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.Operations = append(p.Operations, "sa")
}

func (p *Piles) sb() {
	if len(p.B) < 2 {
		return
	}
	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "sb")
}

func (p *Piles) pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

func (p *Piles) pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

func (p *Piles) rra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	p.Operations = append(p.Operations, "rra")
}

func (p *Piles) rrb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	p.Operations = append(p.Operations, "rrb")
}
