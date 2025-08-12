package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range Commands {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	return nil
}
