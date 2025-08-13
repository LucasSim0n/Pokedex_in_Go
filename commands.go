package main

type cliCommand struct {
	Name        string
	Description string
	Callback    func(c *config, name string) error
}

var Commands map[string]cliCommand

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "bmap",
			Description: "Display the previous 20 locations",
			Callback:    commandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Display pokemons from the selected area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch the selected pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Show data from the selected pokemon",
			Callback:    commandInspect,
		},
	}
}
