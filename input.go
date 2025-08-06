package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Split(lower, " ")
	var result []string
	for _, w := range split {
		fmt.Println(w)
		if strings.TrimSpace(w) != "" {
			result = append(result, w)
		}
	}
	return result
}
