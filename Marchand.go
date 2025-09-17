package main

import (
	"fmt"
)

// Endroit des prix des items
var items = map[string]int{
	"Potion de vie":                     17,
	"Potion de poison":                  6,
	"Boule d'électricité":               2,
	"Armure amovible adaptative":        34,
	"Blaster lazer":                     27,
	"Armes de désintégration intégrale": 200,
	"Plume d’oiseaux célestes":          20,
	"Pistolet Nucléaire":                80,
}

// Stock initial : 5 exemplaires de chaque item
var stock = map[string]int{
	"Potion de vie":                     5,
	"Potion de poison":                  5,
	"Boule d'électricité":               5,
	"Armure amovible adaptative":        5,
	"Blaster lazer":                     5,
	"Armes de désintégration intégrale": 5,
	"Plume d’oiseaux célestes":          5,
	"Pistolet Nucléaire":                5,
}

// AcheterItem essaie d'acheter un item si le joueur a assez de pièces et si l'objet est en stock
func AcheterItem(inventaire *Inventaire, nomItem string, piecesOr *int) {
	prix, existe := items[nomItem]
	if !existe {
		fmt.Println("Cet objet n'existe pas dans la boutique.")
		return
	}

	if stock[nomItem] <= 0 {
		fmt.Println("Cet objet a été épuisé.")
		return
	}

	if *piecesOr < prix {
		fmt.Println("Pas assez de pièces d’or.")
		return
	}

	if len(inventaire.objets) >= inventaire.limite {
		fmt.Println("Votre inventaire est plein.")
		return
	}

	// Achat réussi
	*piecesOr -= prix
	stock[nomItem]--
	inventaire.objets = append(inventaire.objets, nomItem)

	fmt.Printf("Vous avez acheté : %s\n", nomItem)
	fmt.Printf("Il vous reste %d pièces d’or.\n", *piecesOr)
	fmt.Printf("Stock restant de \"%s\" : %d\n", nomItem, stock[nomItem])
}
func DoesItemExist(object string) bool {
	_, exists := items[object]
	return exists
}
