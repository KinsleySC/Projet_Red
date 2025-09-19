package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
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
	qty, exists := c.Inventory["Potion de soin"]
	if !exists || qty <= 0 {
		fmt.Println("Vous n'avez pas de potion !")
		return
	}
	c.Inventory["Potion de soin"]--
	if c.Inventory["Potion de soin"] == 0 {
		delete(c.Inventory, "Potion de soin")
	}
	c.CurrentHp += 50
	if c.CurrentHp > c.HpMax {
		c.CurrentHp = c.HpMax
	}
	fmt.Printf("Vous avez utilisé une potion de soin. PV actuels : %d / %d\n", c.CurrentHp, c.HpMax)
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
		fmt.Printf("Dégâts de poison : %s - PV actuels %d / %d\n", target.Name, target.CurrentHp, target.HpMax)
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

func totalInventory(c *Character) int {
	total := 0
	for _, qty := range c.Inventory {
		total += qty
	}
	return total
}

func canAddItem(c *Character, qtyToAdd int) bool {
	return totalInventory(c)+qtyToAdd <= 10
}

func characterCreation() Character {
	reader := bufio.NewReader(os.Stdin)
	var name, class string

	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)
		if isAlpha(name) && name != "" {
			name = formatName(name)
			break
		}
		fmt.Println("Nom invalide. Veuillez utiliser uniquement des lettres.")
	}

	for {
		fmt.Print("Choisissez une classe (Humain, Elfe, Nain) : ")
		input, _ := reader.ReadString('\n')
		class = strings.TrimSpace(strings.Title(strings.ToLower(input)))
		if class == "Humain" || class == "Elfe" || class == "Nain" {
			break
		}
		fmt.Println("Classe invalide. Veuillez choisir parmi : Humain, Elfe, Nain.")
	}

	var hpMax int
	switch class {
	case "Humain":
		hpMax = 100
	case "Elfe":
		hpMax = 80
	case "Nain":
		hpMax = 120
	}

	currentHp := hpMax / 2
	genre := "Homme"

	inv := map[string]int{
		"Potion de soin":               3,
		"Potion de poison":             1,
		"Livre de sort : Boule de feu": 1,
	}

	return initCharacter(genre, name, class, 1, hpMax, currentHp, inv)
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	name = strings.ToLower(name)
	return strings.ToUpper(string(name[0])) + name[1:]
}

func main() {
	c1 := characterCreation()
	displayInfo(c1)
	accessInventory(c1)
	c1.CurrentHp = 0
	isDead(&c1)
	takePot(&c1)

	enemy := Enemy{
		Name:      "Gobelin",
		HpMax:     40,
		CurrentHp: 40,
	}

	poisonPot(&c1, enemy)
	useSpellBook(&c1)

	fmt.Println("\nTentative d'ajout d'une potion de soin...")
	if canAddItem(&c1, 1) {
		c1.Inventory["Potion de soin"]++
		fmt.Println("Potion de soin ajoutée à l'inventaire.")
	} else {
		fmt.Println("Inventaire plein !")
	}
}
