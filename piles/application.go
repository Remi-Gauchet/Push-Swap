package piles

import (
	"fmt"
)

// lit un string et l'applique à la structure Piles
// renvoie une erreur si ce n'est pas une opération
func (p *Piles) Application(operation string) error {
	switch operation {
	case "ra":
		p.ra()
	case "rb":
		p.rb()
	case "rr":
		p.rr()
	case "sa":
		p.sa()
	case "sb":
		p.sb()
	case "ss":
		p.ss()
	case "pb":
		p.pb()
	case "pa":
		p.pa()
	case "rra":
		p.rra()
	case "rrb":
		p.rrb()
	case "rrr":
		p.rrr()
	default:
		return fmt.Errorf("%q n'est pas une opération", operation)
	}
	return nil
}
