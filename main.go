package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokeapi "github.com/CamilleOnoda/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient          pokeapi.PokeAPIClient //talking to the PokeAPI client
	nextLocationURL     *string               //where to go for the map command's next page
	previousLocationURL *string               //where to go for the map command's previous page
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the name of the previous 20 location areas in the Pokemon world",
		callback:    commandMapb,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeClient: client,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}
		value, ok := commands[cleanedInput[0]]
		if !ok {
			fmt.Print("Unknown command")
		} else {
			if err := value.callback(cfg); err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message" +
		"\nexit: Exit the Pokedex\n")
	return nil
}

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeClient.GetLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return fmt.Errorf("Error fetching location areas: %w", err)
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.previousLocationURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationsResp, err := cfg.pokeClient.GetLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return fmt.Errorf("Error fetching location areas: %w", err)
	}
	cfg.nextLocationURL = locationsResp.Next
	cfg.previousLocationURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil

}
