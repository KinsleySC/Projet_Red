package main

import (
	"fmt"
	"strings"
)

type Inventaire struct {
	objets []string
	limite int
	ouvert bool
}

func NouveauInventaire(limite int) *Inventaire {
	return &Inventaire{
		objets: []string{},
		limite: limite,
		ouvert: false,
	}
}

func (inv *Inventaire) Ouvrir() {
	if inv.ouvert {
		fmt.Println("L'inventaire est déjà ouvert.")
	} else {
		inv.ouvert = true
		fmt.Println("Inventaire ouvert.")
		inv.Afficher()
	}
}

func (inv *Inventaire) Fermer() {
	if !inv.ouvert {
		fmt.Println("L'inventaire est déjà fermé.")
	} else {
		inv.ouvert = false
		fmt.Println("Inventaire fermé.")
	}
}

func (inv *Inventaire) Ajouter(objet string) {
	if len(inv.objets) >= inv.limite {
		fmt.Println("Inventaire plein !")
		return
	}

	if !DoesItemExist(objet) {
		fmt.Printf("L'objet \"%s\" n'existe pas dans la liste des items disponibles.\n", objet)
		return
	}

	inv.objets = append(inv.objets, objet)
	fmt.Printf("%s ajouté à l'inventaire.\n", objet)
}

func (inv *Inventaire) Supprimer(objet string) {
	for i, o := range inv.objets {
		if strings.ToLower(o) == strings.ToLower(objet) {
			inv.objets = append(inv.objets[:i], inv.objets[i+1:]...)
			fmt.Printf("%s supprimé de l'inventaire.\n", objet)
			return
		}
	}
	fmt.Println("Objet non trouvé dans l'inventaire.")
}

func (inv *Inventaire) Afficher() {
	if !inv.ouvert {
		fmt.Println("L'inventaire est fermé.")
		return
	}
	if len(inv.objets) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Contenu de l'inventaire :")
	for i, obj := range inv.objets {
		fmt.Printf("%d. %s\n", i+1, obj)
	}
}
