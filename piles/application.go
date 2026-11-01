package piles

import (
	"fmt"
)

// Application lit un string et l'applique à la structure Piles si c'est une opération
// sinon elle renvoie une erreur
func (p *Piles) Application(operation string) error {
	switch {
	case operation == "ra":
		p.ra()
	case operation == "rb":
		p.rb()
	case operation == "rr":
		p.rr()
	case operation == "sa":
		p.sa()
	case operation == "sb":
		p.sb()
	case operation == "ss":
		p.ss()
	case operation == "pb":
		p.pb()
	case operation == "pa":
		p.pa()
	case operation == "rra":
		p.rra()
	case operation == "rrb":
		p.rrb()
	case operation == "rrr":
		p.rrr()
	default:
		return fmt.Errorf("%q n'est pas une opération", operation)
	}
	return nil
}
