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
}

type Enemy struct {
	Name      string
	HpMax     int
	CurrentHp int
}
