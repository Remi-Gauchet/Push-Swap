package main

// positionDuMaximumB cherche le plus gros int de la pile B et renvoie son index
func (p *Piles) positionDuMaximumB() int {
	var maximum int = p.b[0]
	var position int = 0
	for i := 0; i < len(p.b); i++ {
		if p.b[i] > maximum {
			maximum = p.b[i]
			position = i
		}
	}
	return position
}
