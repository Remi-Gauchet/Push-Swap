package piles

// positionDuMaximumB cherche le plus gros int de la pile B et renvoie son index
func (p *Piles) positionDuMaximumB() int {
	var maximum int = p.B[0]
	var position int = 0
	for i := 0; i < len(p.B); i++ {
		if p.B[i] > maximum {
			maximum = p.B[i]
			position = i
		}
	}
	return position
}
