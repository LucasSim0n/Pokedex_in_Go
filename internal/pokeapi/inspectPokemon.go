package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) InspectPokemon(name string) (*PokemonResStruct, error) {
	var pokemon *PokemonResStruct
	fullURL := baseURL + "pokemon/" + name

	gotPokemon := c.cache.GetCP(fullURL)
	if !gotPokemon {
		return pokemon, fmt.Errorf("you have not caught that pokemon")
	}

	data, ok := c.cache.Get(fullURL)
	if !ok {
		return pokemon, fmt.Errorf("error getting data from cache")
	}

	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("Error decoding json data: %v", err)
	}

	return pokemon, nil
}
