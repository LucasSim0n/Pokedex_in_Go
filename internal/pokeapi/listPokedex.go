package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListPokemon() error {
	poks := c.cache.GetAllPokemons()
	var pokemon PokemonResStruct

	for _, pok := range poks {

		err := json.Unmarshal(pok, &pokemon)
		if err != nil {
			return fmt.Errorf("Error decoding json: %v", err)
		}
		fmt.Println(pokemon.Name)
	}

	return nil
}
