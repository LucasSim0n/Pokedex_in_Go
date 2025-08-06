package main

import (
	"bufio"
	"fmt"
	"github.com/LucasSim0n/Pokedex_in_Go/internal"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := &internal.Config{
		Next: 1,
		Prev: -19,
	}

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		stripText := strings.TrimSpace(scanner.Text())
		input := strings.Split(strings.ToLower(stripText), " ")[0]

		com, ok := internal.Commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := com.Callback(conf)
		if err != nil {
			fmt.Printf("Error occured with command %s: %v\n", com.Name, err)
		}
	}
}
