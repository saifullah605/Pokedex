package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!!")
}

// testing
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
