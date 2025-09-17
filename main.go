package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inv := NouveauInventaire(10)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bienvenue dans le jeu ! Tape 'help' pour les commandes.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		args := strings.Split(input, " ")
		cmd := strings.ToLower(args[0])

		switch cmd {
		case "help":
			fmt.Println("Commandes : open (ouvrir inventaire), close (fermer inventaire), acheter <objet>, jetter, show (afficher), quit (quitter)")
		case "open":
			inv.Ouvrir()
		case "close":
			inv.Fermer()
		case "acheter":
			if len(args) < 2 {
				fmt.Println("Usage : acheter <objet>")
				continue
			}
			objet := strings.Join(args[1:], " ")
			inv.Ajouter(objet)
		case "jetter":
			if len(args) < 2 {
				fmt.Println("Usage : jetter <objet>")
				continue
			}
			objet := strings.Join(args[1:], " ")
			inv.Supprimer(objet)
		case "show":
			inv.Afficher()
		case "quit":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Commande inconnue, tapez 'help' pour la liste.")
		}
	}
}
