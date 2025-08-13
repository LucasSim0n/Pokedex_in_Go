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
		fmt.Print("Pokedex~> ")
		scanner.Scan()
		stripText := strings.TrimSpace(scanner.Text())
		input := strings.Split(strings.ToLower(stripText), " ")
		userCom := input[0]
		locName := ""
		if len(input) > 1 {
			locName = input[1]
		}

		com, ok := Commands[userCom]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := com.Callback(cfg, locName)
		if err != nil {
			fmt.Printf("Error occured with command %s: %v\n", com.Name, err)
		}
	}
}
