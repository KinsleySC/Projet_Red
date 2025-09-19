package main

import (
	"fmt"
)

type Options struct {
	arg string
	val int
}

func MainMenu() int {
	options := []Options{
		{"Jouer", 1},
		{"Options", 2},
		{"Quitter", 3},
	}
	menu := NewMenu(" ----- StarForge -----")
	for _, opti := range options {
		menu.AddItem(opti.arg, opti.val)
	}
	result, err := menu.Display()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	switch v := result.(type) {
	case int:
		if v == 1 {
			SubStart()
		}
		fmt.Printf("Selected option: %d\n", v)
	case string:
		fmt.Printf("Selected option: %s\n", result)
	default:
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result, result)
	}
	return 0
}

func SubStart() {
	option := []Options{
		{"Nouvelle partie", 1},
		{"Retour", 2},
	}
	menu2 := NewMenu("Partie")
	for _, opti := range option {
		menu2.AddItem(opti.arg, opti.val)
	}
	result1, err1 := menu2.Display()
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	}
	switch v := result1.(type) {
	case int:
		fmt.Printf("Selectionner l'option: %d\n\n", v)
		if v == 2 {
			MainMenu()
		}
		if v == 1 {
			Start_CreationOfCharacter()
		}
	case string:
		fmt.Printf("Selected option: %s\n\n", v)
	default:
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result1, result1)
	}
}

//func SubOption() {
//
//	menu2 := NewMenu("Options")
//	if option == "2.Option" {
//		menu2.AddItem("Music", 1)
//		menu2.AddItem("Back", 2)
//		result1, err1 := menu2.Display()
//		if err1 != nil {
//			fmt.Printf("Error: %v\n", err1)
//		}
//		if _, ok := result1.(int); ok {
//			fmt.Printf("Selected option: %d\n", result1)
//		} else if _, ok := result1.(string); ok {
//			fmt.Printf("Selected option: %s\n", result1)
//		} else {
//			fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result1, result1)
//		}
//	}
//}
