package internal

import (
	"fmt"
	"os"
)

type Config struct {
	Next int
	Prev int
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range Commands {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	return nil
}

func commandMap(c *Config) error {

	for i := range 20 {
		id := c.Next + i

		nLoc, err := GetLocation(id)
		if err != nil {
			return err
		}

		fmt.Println(nLoc)

	}
	c.Next += 20
	c.Prev += 20
	return nil
}

func commandMapB(c *Config) error {

	if c.Prev < 1 {
		return fmt.Errorf("you're on the first page")
	}

	for i := range 20 {
		id := c.Prev + i

		nLoc, err := GetLocation(id)
		if err != nil {
			return err
		}

		fmt.Println(nLoc)
	}
	c.Prev -= 20
	c.Next -= 20

	return nil
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
	}
}
