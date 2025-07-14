package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	pokeapi "github.com/saifullah605/Pokedex/PokeAPI"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var Pokedex = map[string]pokeapi.Pokemon{}

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
	if len(cleaned) == 1 {
		return fmt.Errorf("invalid entry")
	}

	hyphenated := stringHyphenated(cleaned[1:])
	areaData, err := pokeapi.ExploreRequest(hyphenated)

	if err != nil {
		return err
	}

	fmt.Println("Exploring " + hyphenated + "...")

	for _, encounter := range areaData.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch() error {
	if len(cleaned) == 1 {
		return fmt.Errorf("invalid entry")
	}

	pokemon, err := pokeapi.PokemonRequest(cleaned[1])

	if err != nil {
		return err
	}

	fmt.Print("Throwing a Pokeball at " + cleaned[1])

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	time.Sleep(1 * time.Second)
	fmt.Print("\n")

	if isCaught(pokemon.BaseExperience) {
		Pokedex[cleaned[1]] = pokemon
		fmt.Println(cleaned[1] + " was caught!")
	} else {
		fmt.Println(cleaned[1] + " escaped!")
	}

	return nil

}

func commandInspect() error {
	if len(cleaned) == 1 {
		return fmt.Errorf("invalid entry")
	}

	pokemon, valid := Pokedex[cleaned[1]]

	if !valid {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name: " + cleaned[1])
	fmt.Println("Height:", pokemon.Height, "\nWeight:", pokemon.Weight)
	fmt.Println("Stats")
	for _, stat := range pokemon.Stats {
		fmt.Println("\t-", stat.Stat.Name+":", stat.BaseStat)
	}

	fmt.Println("Types: ")
	for _, Type := range pokemon.Types {
		fmt.Println("\t-", Type.Type.Name)
	}

	return nil
}

func commandPokedex() error {
	fmt.Println("Your Pokedex:")

	for name := range Pokedex {
		fmt.Println("\t-", name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit":    {"exit", "Exit the Pokedex", commandExit},
		"help":    {"help", "Displays a help message", commandHelp},
		"map":     {"map", "Display 20 loactions", commandMap},
		"mapb":    {"mapb", "Displays the previous 20 locations", commandMapb},
		"explore": {"explore", "Display diffrent pokemon with a area input example: explore canalave city", commandExplore},
		"catch":   {"catch", "Try to catch a pokemon, example: catch pidgey", commandCatch},
		"inspect": {"inspect", "inspect the caught pokemon stats", commandInspect},
		"pokedex": {"pokedex", "list all caught pokemon", commandPokedex},
	}
}

func stringHyphenated(words []string) string {
	properString := ""
	for i, word := range words {
		if i < len(words)-1 {
			properString += word + "-"
		} else {
			properString += word
		}

	}
	return properString

}

func isCaught(baseExp int) bool {
	threshold := 25
	r := rand.Intn(baseExp + 1)

	return r < threshold
}
