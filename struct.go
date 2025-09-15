package piscine

type Monster struct {
	name         string
	hp_max       int
	current_hp   int
	attack_point int
}

type Character struct {
	genre      string
	name       string
	classe     string
	level      int
	hp_max     int
	current_hp int
	inventory  map[string]int
}
