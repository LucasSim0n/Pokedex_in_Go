package main

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

func GetLocation(id int) (string, error) {
	var data Location

	fullURL := "https://pokeapi.co/api/v2/location-area/" + strconv.Itoa(id)
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

	return data.Name, nil
}
