package main

import (
	"fmt"
	"os"
)

type config struct {
	next int
	prev int
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range Commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c *config) error {

	for i := range 20 {
		id := c.next + i

		nLoc, err := GetLocation(id)
		if err != nil {
			return err
		}

		fmt.Println(nLoc)

	}
	c.next += 20
	c.prev += 20
	return nil
}

func commandMapB(c *config) error {

	if c.prev < 1 {
		return fmt.Errorf("you're on the first page")
	}

	for i := range 20 {
		id := c.prev + i

		nLoc, err := GetLocation(id)
		if err != nil {
			return err
		}

		fmt.Println(nLoc)
	}
	c.prev -= 20
	c.next -= 20

	return nil
}

var Commands map[string]cliCommand

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "bmap",
			description: "Display the previous 20 locations",
			callback:    commandMapB,
		},
	}
}
