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

func accessInventory(c Character) {
	fmt.Println("Objets dans l'inventaire :")
	for item, qty := range c.Inventory {
		fmt.Printf("%s : %d\n", item, qty)
	}
}

func takePot(c *Character) {
	qty, exists := c.Inventory["Potion"]
	if !exists || qty <= 0 {
		fmt.Println("Vous n'avez pas de potion !")
		return
	}
	c.Inventory["Potion"]--
	if c.Inventory["Potion"] == 0 {
		delete(c.Inventory, "Potion")
	}
	c.CurrentHp += 50
	if c.CurrentHp > c.HpMax {
		c.CurrentHp = c.HpMax
	}

	fmt.Printf("Vous avez utilisé une potion. PV actuels : %d / %d\n", c.CurrentHp, c.HpMax)
}

func isDead(c *Character) {
	if c.CurrentHp <= 0 {
		fmt.Println("Vous êtes mort... ")
		c.CurrentHp = c.HpMax / 2
		fmt.Printf("Vous êtes ressuscité avec %d / %d PV.\n", c.CurrentHp, c.HpMax)
	}
}

func main() {
	var playerName string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&playerName)

	inv := map[string]int{"Potion": 3}

	c1 := initCharacter("Homme", playerName, "Elfe", 1, 100, 40, inv)

	displayInfo(c1)

	accessInventory(c1)

	c1.CurrentHp = 0
	isDead(&c1)
}
