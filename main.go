package main

import "fmt"

func main() {
	menu := NewMenu("Game")
	menu.AddItem("Start", 1)
	menu2 := NewMenu("Party")
	menu2.AddItem("New game", 1)
	menu.AddItem("Option", 2)
	menu.AddItem("Exit", 3)

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
	result1, err1 := menu2.Display()
	if err1 != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if _, ok := result1.(int); ok {
		fmt.Printf("Selected option: %d\n", result1)
	} else if _, ok := result1.(string); ok {
		fmt.Printf("Selected option: %s\n", result1)
	} else {
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result1, result1)
	}
}
