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

	fmt.Println("Bienvenue dans le jeu ! Tape 'commandes' pour afficher les commandes.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		args := strings.Split(input, " ")
		cmd := strings.ToLower(args[0])

		switch cmd {
		case "commandes":
			fmt.Println("Commandes : open (ouvrir inventaire), close (fermer inventaire),boutique, jetter, afficher, quitter, banque")
		case "open":
			inv.Ouvrir()
		case "close":
			inv.Fermer()
		case "boutique":
			for k, v := range Items {
				fmt.Printf("%s: %d\n", k, v)
			}
		case "banque":
			{
				print("bienvenu dans la banque vous possedez <argent")
			}
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
		case "afficher":
			inv.Afficher()
		case "quitter":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Commande inconnue, tapez 'commandes' pour la liste.")
		}
	}
}
