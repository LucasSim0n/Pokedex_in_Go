package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		stripText := strings.TrimSpace(scanner.Text())
		input := strings.Split(strings.ToLower(stripText), " ")[0]

		c, ok := Commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := c.callback()
		if err != nil {
			fmt.Printf("Error occured with command %s: %v", c.name, err)
		}
	}
}
