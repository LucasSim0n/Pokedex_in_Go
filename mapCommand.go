package main

import (
	"fmt"
	"github.com/LucasSim0n/Pokedex_in_Go/internal/pokeapi"
)

func commandMap(c *config, name string) error {

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

func commandMapB(c *config, name string) error {

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
