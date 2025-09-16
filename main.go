package main

import (
	"fmt"
)

func main() {

	menu := NewMenu("Chose a colour")

	menu.AddItem("Langues", 1)
	menu.AddItem("Commandes", 2)
	menu.AddItem("Musiques", 3)
	menu.AddItem("Star select niveau", 4)
	menu.AddItem("Jouer", 5)
	menu.AddItem("Retour", 6)
	menu.AddItem("Quiter", 7)
	menu.AddItem("Yellow", 8)

	result, err := menu.Display()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if _, ok := result.(int); ok {
		fmt.Printf("Selected option: %d\n", result)
	} else if _, ok := result.(string); ok {
		fmt.Printf("Selected option: %s\n", result)
	} else {
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result, result)
	}
}
