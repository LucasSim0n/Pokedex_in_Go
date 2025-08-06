package main

import (
	"github.com/LucasSim0n/Pokedex_in_Go/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		apiClient: pokeClient,
		next:      1,
		prev:      -19,
	}

	startRepl(cfg)
}
