package main

import (
	"bufio"
	"fmt"
	"github.com/LucasSim0n/Pokedex_in_Go/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	next      int
	prev      int
	apiClient pokeapi.Client
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		stripText := strings.TrimSpace(scanner.Text())
		input := strings.Split(strings.ToLower(stripText), " ")[0]

		com, ok := Commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := com.Callback(cfg)
		if err != nil {
			fmt.Printf("Error occured with command %s: %v\n", com.Name, err)
		}
	}
}
