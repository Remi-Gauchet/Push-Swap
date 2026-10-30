package main

import (
	"fmt"
	"os"

	"github.com/Remi-Gauchet/Algorithmie/seance5/outils"
)

// structure
type Piles struct {
	a          []int
	b          []int
	operations []string
}

// exécute les fonctions et renvoie une erreur
func executer(args []string) error {
	if len(os.Args) < 2 {
		return nil
	}

	//appel de la fonction comparer
	if args[0] == "comparer" {
		comparer()
		return nil
	}

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

	p := Piles{a: outils.Rangs(nombres)}
	// une pile déjà triée est renvoyée immédiatement
	if p.aEstTriee() {
		return nil
	}

	// on commence le tri
	p.trier()

	//si le programme tel quel n'arrive pas à trier
	if p.aEstTriee() == false || len(p.b) > 0 {
		return fmt.Errorf("impossible de trier la pile")
	}

	// déclare les opérations effectuées pour le tri
	fmt.Println("opérations :", p.operations)
	fmt.Println("nombre d'opérations :", len(p.operations))
	return nil
}

// Lance le programme et l'arrête en renvoyant une erreur si besoin
func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}
