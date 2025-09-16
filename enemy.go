package piscine

func initGoblin() {
	monster := Monster{"Gobelin d'entraînement", 40, 40, 5}
}

func goblinPattern() {
	for i := 0; i < 3; i++ {
		initGoblin()
	}
}
