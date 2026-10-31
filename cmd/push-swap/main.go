package main

import (
	"fmt"
	"os"

	"github.com/Remi-Gauchet/Push-Swap/outils"
	"github.com/Remi-Gauchet/Push-Swap/piles"
)

// exécute les fonctions et renvoie une erreur
func executer(args []string) error {
	//on lit l'argument
	nombres, err := outils.LireNombres(args)
	if err != nil {
		return err
	}

	//on vérifie qu'il y ait bien un argument
	if len(nombres) == 0 {
		return nil
	}

	//on n'accepte pas les doublons
	err = outils.VerifierDoublons(nombres)
	if err != nil {
		return err
	}

	p := piles.Piles{A: outils.Rangs(nombres)}
	// une pile déjà triée est renvoyée immédiatement
	if p.AEstTriee() {
		return nil
	}

	// on commence le tri
	p.Trier()

	//si le programme tel quel n'arrive pas à trier
	if p.AEstTriee() == false || len(p.B) > 0 {
		return fmt.Errorf("impossible de trier la pile")
	}

	// déclare les opérations effectuées pour le tri, avec une opération par ligne
	for _, op := range p.Operations {
		fmt.Println(op)
	}
	return nil
}

// Lance le programme et l'arrête en renvoyant une erreur si besoin
func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
}
