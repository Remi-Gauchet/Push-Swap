package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Remi-Gauchet/Push-Swap/outils"
	"github.com/Remi-Gauchet/Push-Swap/piles"
)

// executer lit les opérations qu'on lui donne et réorganise la structure Piles en fonction
// à la fin, il renvoie OK si Piles est bien triée, et KO elle ne l'est pas
func executer(args []string) error {
	// on lit l'argument
	nombres, err := outils.LireNombres(args)
	if err != nil {
		return err
	}

	// on vérifie qu'il y ait bien un argument
	if len(nombres) == 0 {
		return nil
	}

	// on n'accepte pas les doublons
	err = outils.VerifierDoublons(nombres)
	if err != nil {
		return err
	}

	// on définit la pile
	p := piles.Piles{A: nombres}

	// on scanne les lignes d'opérations
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() == true {
		ligne := scanner.Text()
		// on appelle la fonction Application, qui modifie les Piles selon les operations lues
		err = p.Application(ligne)
		if err != nil {
			return err
		}
	}
	// on renvoie une erreur s'il y a un problème de lecture
	err = scanner.Err()
	if err != nil {
		return err
	}

	// on écrit le retour de Checker selon si la structure Piles est triée ou non
	if p.AEstTriee() == false || len(p.B) > 0 {
		fmt.Println("KO")
	} else {
		fmt.Println("OK")
	}
	return nil
}

// main du checker appelle la fonction executer, et renvoie une erreur si besoin
func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
}
