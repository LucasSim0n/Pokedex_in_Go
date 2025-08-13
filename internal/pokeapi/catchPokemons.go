package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (bool, error) {
	xp, err := getPokemonXP(c, name)
	if err != nil {
		return false, err
	}

	diff := GetDifficulty(xp)
	if diff.Chance > rand.Intn(100) {
		return true, nil
	}

	return false, nil
}

func getPokemonXP(c *Client, name string) (int, error) {
	fullURL := baseURL + "pokemon/" + name
	var pokemon PokemonResStruct

	if data, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return 0, fmt.Errorf("Error decoding cached data: %v", err)
		}
		return pokemon.BaseExperience, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return 0, fmt.Errorf("Error with http request: %v", err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("Error reading response data: %v", err)
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return 0, fmt.Errorf("Error decoding json data: %v", err)
	}

	c.cache.Add(fullURL, data)

	return pokemon.BaseExperience, nil

}
