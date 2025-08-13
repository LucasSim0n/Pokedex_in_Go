package main

import (
	"fmt"
)

func commandHelp(c *config, name string) error {
	fmt.Println("\n[+] Welcome to the Pokedex!\n\n\tUsage:\n")
	for _, c := range Commands {
		fmt.Printf("[-] %s: %s\n\n", c.Name, c.Description)
	}
	return nil
}
