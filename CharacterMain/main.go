package main

import "fmt"

func initCharacter(genre string, name string, classe string, level int, hp_max int, current_hp int, inventory map[string]int) Character {
	return Character{
		genre:      genre,
		name:       name,
		classe:     classe,
		level:      level,
		hp_max:     hp_max,
		current_hp: current_hp,
		inventory:  inventory,
	}
}

func main() {
	inv := map[string]int{
		"Potion": 3,
	}
	c1 := initCharacter("Homme", "Nom", "Elfe", 1, 100, 40, inv)

	fmt.Println("Personnage créé :")
	fmt.Println("Nom :", c1.name)
	fmt.Println("Classe :", c1.classe)
	fmt.Println("Niveau :", c1.level)
	fmt.Println("PV :", c1.current_hp, "/", c1.hp_max)
	fmt.Println("Inventaire :", c1.inventory)
}
