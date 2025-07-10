package main

import (
	"fmt"
	"os"

	pokeapi "github.com/saifullah605/Pokedex/PokeAPI"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil

}

func commandMap() error {
	locations, err := pokeapi.MapRequest()
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil

}

func commandMapb() error {
	locations, err := pokeapi.PrevMapRequest()

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore() error {
	fmt.Println(cleaned)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit":    {"exit", "Exit the Pokedex", commandExit},
		"help":    {"help", "Displays a help message", commandHelp},
		"map":     {"map", "Display 20 loactions", commandMap},
		"mapb":    {"mapb", "Displays the previous 20 locations", commandMapb},
		"explore": {"explore", "Display diffrent pokemon with a area input example: explore canalave city", commandExplore},
	}
}
