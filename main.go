package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	pokeapi "github.com/CamilleOnoda/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient          pokeapi.PokeAPIClient
	nextLocationURL     *string
	previousLocationURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
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
	"explore": {
		name:        "explore",
		description: "See a list of all the Pokemon located in a specific location area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a Pokemon and add it to your Pokedex",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "View detailed information about a specific Pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "View all the Pokemon you have caught so far",
		callback:    commandPokedex,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeClient:    client,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
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
			fmt.Print("Unknown command\n\n")
		} else {
			arg := ""
			if len(cleanedInput) > 1 {
				arg = cleanedInput[1]
			}
			if err := value.callback(cfg, arg); err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		}
	}
}

func commandExit(cfg *config, args string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message" +
		"\nexit: Exit the Pokedex\n")
	return nil
}

func commandMap(cfg *config, args string) error {
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

func commandMapb(cfg *config, args string) error {
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

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		return fmt.Errorf("Please provide a location to explore")
	}

	fmt.Printf("Exploring %s...\n\n", areaName)
	pokemonResp, err := cfg.pokeClient.GetPokemonNamesInLocationArea(areaName)
	if err != nil {
		return fmt.Errorf("Error fetching Pokemon in location area '%s': %w", areaName, err)
	}

	if len(pokemonResp.PokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in location area '%s'.\n", areaName)
		return nil
	}

	fmt.Printf("\nPokemon found:\n")
	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		fmt.Printf("Please provide the name of the Pokemon to catch\n\n")
		return nil
	}

	userBaseExperience := rand.Intn(201) + 50
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return fmt.Errorf("Error fetching Pokemon '%s': %w", pokemonName, err)
	}

	if userBaseExperience > pokemon.BaseExperience {
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n\n", pokemonName)

	} else {
		fmt.Printf("%s escaped!\n\n", pokemonName)
	}

	return nil
}

func commandInspect(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		fmt.Printf("Please provide the name of the Pokemon to inspect\n\n")
		return nil
	}
	if cfg.caughtPokemon[pokemonName].Name == "" {
		fmt.Printf("You haven't caught '%s' yet\n\n", pokemonName)
		return nil
	}

	fmt.Printf("Inspecting %s...\n\n", pokemonName)

	pokemon, err := cfg.pokeClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return fmt.Errorf("Error fetching Pokemon '%s': %w", pokemonName, err)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("- %s\n", typeInfo.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Printf("You haven't caught any Pokemon yet\n\n")
		return nil
	}

	fmt.Printf("Your Pokedex:\n\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n\n", pokemon.Name)
	}
	return nil
}
