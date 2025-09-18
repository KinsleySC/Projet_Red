package main

type Character struct {
	Genre     string
	Name      string
	Classe    string
	Level     int
	HpMax     int
	CurrentHp int
	Inventory map[string]int
	Skills    []string
	Equipment
}

type Enemy struct {
	Name      string
	HpMax     int
	CurrentHp int
}

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}
