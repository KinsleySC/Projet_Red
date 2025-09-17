package main

import (
	"fmt"
	"time"
)

func initCharacter(genre, name, classe string, level, hpMax, currentHp int, inventory map[string]int) Character {
	return Character{
		Genre:     genre,
		Name:      name,
		Classe:    classe,
		Level:     level,
		HpMax:     hpMax,
		CurrentHp: currentHp,
		Inventory: inventory,
		Skills:    []string{"Coup de poing"},
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
	fmt.Println("Compétences :", c.Skills)
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

func poisonPot(c *Character, target Enemy) {
	qty, exists := c.Inventory["Potion de poison"]
	if !exists || qty <= 0 {
		fmt.Println("Vous n'avez pas de potion de poison !")
		return
	}
	c.Inventory["Potion de poison"]--
	if c.Inventory["Potion de poison"] == 0 {
		delete(c.Inventory, "Potion de poison")
	}

	fmt.Println("Vous avez utilisé une potion de poison.")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		target.CurrentHp -= 10
		if target.CurrentHp < 0 {
			target.CurrentHp = 0
		}
		fmt.Printf("Dégâts de poison : %s - PV actuels %d / %d\n",
			target.Name, target.CurrentHp, target.HpMax)
	}

	if target.CurrentHp == 0 {
		fmt.Printf("%s est mort empoisonné !\n", target.Name)
	}

	isDead(c)
}

func spellBook(c *Character, spell string) {
	for _, s := range c.Skills {
		if s == spell {
			fmt.Printf("Le sort %s est déjà appris !\n", spell)
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	fmt.Printf("Nouveau sort appris : %s\n", spell)
}

func useSpellBook(c *Character) {
	qty, exists := c.Inventory["Livre de sort : Boule de feu"]
	if !exists || qty <= 0 {
		fmt.Println("Vous n'avez pas de livre de sort : Boule de feu")
		return
	}
	c.Inventory["Livre de sort : Boule de feu"]--
	if c.Inventory["Livre de sort : Boule de feu"] == 0 {
		delete(c.Inventory, "Livre de sort : Boule de feu")
	}

	spellBook(c, "Boule de feu")
}

func main() {
	var playerName string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&playerName)

	inv := map[string]int{
		"Potion":                       3,
		"Potion de poison":             1,
		"Livre de sort : Boule de feu": 1,
	}

	c1 := initCharacter("Homme", playerName, "Elfe", 1, 100, 40, inv)

	displayInfo(c1)

	accessInventory(c1)

	c1.CurrentHp = 0
	isDead(&c1)

	takePot(&c1)

	enemy := Enemy{
		Name:      "Gobelin",
		HpMax:     50,
		CurrentHp: 50,
	}

	poisonPot(&c1, enemy)

	useSpellBook(&c1)
}
