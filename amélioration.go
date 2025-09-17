package main

import (
	"fmt"
)

// Endroit des prix des items
var items = map[string]int{
	"Potion de vie":                  17,
	"Potion de poison":               6,
	"Boule d’électricité":            25,
	"Armure amovible adaptative":     4,
	"Blaster lazer":                  7,
	"Armes de désintégration intégrale": 200,
	"Plume d’oiseaux célestes":       1,
}

// Stock initial
var stock = map[string]int{
	"Potion de vie":                  5,
	"Potion de poison":               5,
	"Boule d’électricité":            5,
	"Armure amovible adaptative":     5,
	"Blaster lazer":                  5,
	"Armes de désintégration intégrale": 5,
	"Plume d’oiseaux célestes":       5,
}

// Fonction pour améliorer un objet
func ameliorerItem(nom string, valeur int) {
	if _, existe := items[nom]; existe {
		items[nom] += valeur
		fmt.Printf("L'objet \"%s\" a été amélioré ! Nouvelle valeur : %d\n", nom, items[nom])
	} else {
		fmt.Printf("⚠️ L'objet \"%s\" n'existe pas.\n", nom)
	}
}

func main() {
	// Exemple d'amélioration
	ameliorerItem("Potion de vie", 5)
	ameliorerItem("Blaster lazer", 10)
	ameliorerItem("Épée magique", 3) // objet inexistant
}
