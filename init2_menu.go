package main

import "fmt"

func MainMenu() int {
	menu := NewMenu(" ----- Game -----")
	menu.AddItem("1.Start", 1)
	menu.AddItem("2.Option", 2)
	menu.AddItem("3.Exit", 3)

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
	return 0
}

func SubStart() {
	var start string

	menu2 := NewMenu("Party")
	if start == "1.Start" {
		menu2.AddItem("New game", 1)
		menu2.AddItem("Back", 2)
		result1, err1 := menu2.Display()
		//if menu2.AddItem("Back") == "Back" {
		//	return MainMenu()
		//}
		if err1 != nil {
			fmt.Printf("Error: %v\n", err1)
		}
		if _, ok := result1.(int); ok {
			fmt.Printf("Selected option: %d\n", result1)
		} else if _, ok := result1.(string); ok {
			fmt.Printf("Selected option: %s\n", result1)
		} else {
			fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result1, result1)
		}
	} else {
		//function for other verification options
	}
}

func SubOption() {
	var option string

	menu2 := NewMenu("Options")
	if option == "2.Option" {
		menu2.AddItem("Music", 1)
		menu2.AddItem("Back", 2)
		result1, err1 := menu2.Display()
		if err1 != nil {
			fmt.Printf("Error: %v\n", err1)
		}
		if _, ok := result1.(int); ok {
			fmt.Printf("Selected option: %d\n", result1)
		} else if _, ok := result1.(string); ok {
			fmt.Printf("Selected option: %s\n", result1)
		} else {
			fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result1, result1)
		}
	}
}
