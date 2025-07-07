package main

import (
	"fmt"
	"os"
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {"exit", "Exit the Pokedex", commandExit},
		"help": {"help", "Displays a help message", commandHelp},
	}
}
