package main

import "fmt"

func initCharacter(genre, name, classe string, level, hpMax, currentHp int, inventory map[string]int) Character {
	return Character{
		Genre:     genre,
		Name:      name,
		Classe:    classe,
		Level:     level,
		HpMax:     hpMax,
		CurrentHp: currentHp,
		Inventory: inventory,
	}
}

func displayInfo(c Character) {
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Level)
	fmt.Printf("PV : %d / %d\n", c.CurrentHp, c.HpMax)
	fmt.Println("Inventaire :")
	for item, qty := range c.Inventory {
		fmt.Printf("%s : %d\n", item, qty)
	}
}

func main() {
	var playerName string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&playerName)

	inv := map[string]int{"Potion": 3}

	c1 := initCharacter("Homme", playerName, "Elfe", 1, 100, 40, inv)

	displayInfo(c1)
}
