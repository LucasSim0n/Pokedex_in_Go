package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Location struct {
	Name string `json:"name"`
}

func (c *Client) GetLocation(id int) (string, error) {
	var data Location
	fullURL := LocURL + strconv.Itoa(id)

	if val, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return "", fmt.Errorf("Error decoding from cache: %v", err)
		}
		return data.Name, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response: %v", err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", fmt.Errorf("Error decoding json: %v", err)
	}

	c.cache.Add(fullURL, body)
	return data.Name, nil
}
