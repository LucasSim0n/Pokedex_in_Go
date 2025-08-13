package main

import "fmt"

func commandCatch(cfg *config, name string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	success, err := cfg.apiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	if success {
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
