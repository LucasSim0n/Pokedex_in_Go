package main

import (
	"fmt"
	"os"

	"github.com/LucasSim0n/Pokedex_in_Go/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(c *config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range Commands {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	return nil
}

func commandMap(c *config) error {

	for i := range pokeapi.NumLocations {
		id := c.next + i

		nLoc, err := c.apiClient.GetLocation(id)
		if err != nil {
			return err
		}

		fmt.Println(nLoc)

	}
	c.next += pokeapi.NumLocations
	c.prev += pokeapi.NumLocations
	return nil
}

func commandMapB(c *config) error {

	if c.prev < 1 {
		return fmt.Errorf("you're on the first page")
	}

	for i := range pokeapi.NumLocations {
		id := c.prev + i

		nLoc, err := c.apiClient.GetLocation(id)
		if err != nil {
			return err
		}

		fmt.Println(nLoc)
	}
	c.prev -= pokeapi.NumLocations
	c.next -= pokeapi.NumLocations

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
