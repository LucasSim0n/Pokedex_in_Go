package main

import "fmt"

var Stats = map[string]bool{
	"hp":              true,
	"attack":          true,
	"defense":         true,
	"special-attack":  true,
	"special-defense": true,
	"speed":           true,
}

func commandInspect(cfg *config, name string) error {

	pokemon, err := cfg.apiClient.InspectPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemon.Stats {
		if ok := Stats[val.Stat.Name]; ok {
			fmt.Printf("-%s: %d\n", val.Stat.Name, val.BaseStat)
		}
	}
	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("-%s\n", val.Type.Name)
	}

	return nil
}
