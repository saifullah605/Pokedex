package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//pokeapi "github.com/saifullah605/Pokedex/PokeAPI"
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
		cmd, ok := getCommands()[cleaned[0]]

		if !ok {
			fmt.Println("Unknown Command")
			continue
		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}

}
