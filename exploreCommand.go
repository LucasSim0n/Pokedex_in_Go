package main

import "fmt"

func commandExplore(cfg *config, name string) error {
	poks, err := cfg.apiClient.ExplorePokemons(name)
	if err != nil {
		return err
	}
	for _, pok := range poks {
		fmt.Println(pok)
	}
	return nil
}
