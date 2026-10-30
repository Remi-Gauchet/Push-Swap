package main

// positionDuMinimum cherche le plus petit int de la pile A et renvoie son index
func (p *Piles) positionDuMinimum() int {
	var minimum int = p.a[0]
	var position int = 0
	for i := 0; i < len(p.a); i++ {
		if p.a[i] < minimum {
			minimum = p.a[i]
			position = i
		}
	}
	return position
}
