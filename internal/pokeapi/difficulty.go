package pokeapi

type Difficulty struct {
	Name   string
	MinXP  int
	Chance int
}

var difficulties = []Difficulty{
	{"VERY_HARD", 255, 5},
	{"HARD", 200, 16},
	{"MID", 150, 37},
	{"MID_EASY", 100, 63},
	{"EASY", 50, 85},
	{"VERY_EASY", 20, 90},
}

func GetDifficulty(xp int) Difficulty {
	for _, d := range difficulties {
		if xp >= d.MinXP {
			return d
		}
	}
	return difficulties[len(difficulties)-1]
}
