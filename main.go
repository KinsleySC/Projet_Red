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

func main() {
	var playerName string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&playerName)

	inv := map[string]int{"Potion": 3}

	c1 := initCharacter("Homme", playerName, "Elfe", 1, 100, 40, inv)

	fmt.Println("Nom :", c1.Name)
	fmt.Println("Classe :", c1.Classe)
	fmt.Println("Niveau :", c1.Level)
	fmt.Println("PV :", c1.CurrentHp, "/", c1.HpMax)
	fmt.Println("Inventaire :")
	for item, qty := range c1.Inventory {
		fmt.Printf("%s : %d\n", item, qty)
	}

}
