package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	replStart()

}

// testing
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func replStart() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Invalid entry")
			continue
		}

		input := scanner.Text()

		cleaned := cleanInput(input)
		fmt.Println("Your command was: " + cleaned[0])

	}

}
