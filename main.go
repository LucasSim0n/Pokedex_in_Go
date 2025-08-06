package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := &config{
		next: 1,
		prev: -19,
	}

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
		err := com.callback(conf)
		if err != nil {
			fmt.Printf("Error occured with command %s: %v\n", com.name, err)
		}
	}
}
