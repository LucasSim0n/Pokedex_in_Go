package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (bool, error) {
	fullURL := baseURL + "pokemon/" + name
	xp, err := getPokemonXP(c, fullURL)
	if err != nil {
		return false, err
	}

	diff := GetDifficulty(xp)
	if diff.Chance > rand.Intn(100) {
		c.cache.UpdateCP(fullURL)
		return true, nil
	}

	return false, nil
}

func getPokemonXP(c *Client, fullURL string) (int, error) {
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

	c.cache.Add(fullURL, data, false)

	return pokemon.BaseExperience, nil

}
