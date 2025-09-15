package piscine

type Character struct {
	genre      string
	name       string
	classe     string
	level      int
	hp_max     int
	current_hp int
	inventory  map[string]int
}
