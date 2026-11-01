package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Remi-Gauchet/Push-Swap/outils"
	"github.com/Remi-Gauchet/Push-Swap/piles"
)

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

	p := piles.Piles{A: nombres}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() == true {
		ligne := scanner.Text()
		err = p.Application(ligne)
		if err != nil {
			return err
		}
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	if p.AEstTriee() == false || len(p.B) > 0 {
		fmt.Println("KO")
	} else {
		fmt.Println("OK")
	}
	return nil
}

func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
}
