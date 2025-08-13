package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExplorePokemons(name string) ([]string, error) {
	var poks []string

	data, err := getPokemons(c, name)
	if err != nil {
		return poks, err
	}
	for _, encounter := range data.PokemonEncounters {
		poks = append(poks, encounter.Pokemon.Name)
	}

	return poks, nil
}

func getPokemons(c *Client, name string) (LocResStruct, error) {
	var data LocResStruct

	fullURL := baseURL + "location-area/" + name

	if val, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return data, fmt.Errorf("Error decoding cached data: %v", err)
		}
		return data, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return data, fmt.Errorf("Error with http request: %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return data, fmt.Errorf("Error reading response: %v", err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, fmt.Errorf("Error decoding json: %v", err)
	}

	c.cache.Add(fullURL, body)
	return data, nil
}
