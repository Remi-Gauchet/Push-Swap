package piles

// positionDuMinimum cherche le plus petit int de la pile A et renvoie son index
func (p *Piles) positionDuMinimum() int {
	var minimum int = p.A[0]
	var position int = 0
	for i := 0; i < len(p.A); i++ {
		if p.A[i] < minimum {
			minimum = p.A[i]
			position = i
		}
	}
	return position
}
