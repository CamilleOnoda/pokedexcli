package main

import (
	"fmt"
	"testing"

	"github.com/CamilleOnoda/pokedexcli/internal/pokeapi"
)

func TestCaughtCount(t *testing.T) {
	tests := []struct {
		name          string
		catchTimes    int
		expectedCount int
	}{
		{
			name:          "catch once",
			catchTimes:    1,
			expectedCount: 1,
		},
		{
			name:          "catch multiple times",
			catchTimes:    3,
			expectedCount: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := &config{
				caughtPokemon: make(map[string]pokeapi.Pokemon),
				caughtCount:   make(map[string]int),
			}

			pokemon := pokeapi.Pokemon{Name: "pikachu", BaseExperience: 50}

			for i := 0; i < test.catchTimes; i++ {
				if _, exists := cfg.caughtPokemon["pikachu"]; !exists {
					cfg.caughtPokemon["pikachu"] = pokemon
				}
				cfg.caughtCount["pikachu"]++
			}

			if cfg.caughtCount["pikachu"] != test.expectedCount {
				t.Errorf("Expected count %d, got %d", test.expectedCount, cfg.caughtCount["pikachu"])
			}

			if len(cfg.caughtPokemon) != 1 {
				t.Errorf("Expected 1 pokemon in pokedex, got %d", len(cfg.caughtPokemon))
			}
		})
	}
}

func TestCommandMap(t *testing.T) {
	tests := []struct {
		name          string
		mockResponse  pokeapi.LocationAreaResponse
		mockError     error
		expectedError bool
	}{
		{
			name: "returns location areas",
			mockResponse: pokeapi.LocationAreaResponse{
				Next:     nil,
				Previous: nil,
				Results: []pokeapi.LocationArea{
					{Name: "pallet-town"},
					{Name: "viridian-city"},
				},
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "empty results",
			mockResponse:  pokeapi.LocationAreaResponse{},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "API error",
			mockResponse:  pokeapi.LocationAreaResponse{},
			mockError:     fmt.Errorf("API unavailable"),
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := &config{
				pokeClient: &mockClient{
					getLocationAreasFunc: func(pageURL *string) (pokeapi.LocationAreaResponse, error) {
						return test.mockResponse, test.mockError
					},
				},
			}

			err := commandMap(cfg, "")

			if test.expectedError && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !test.expectedError && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}

			if !test.expectedError {
				if cfg.nextLocationURL != test.mockResponse.Next {
					t.Error("Expected nextLocationURL to be updated")
				}
				if cfg.previousLocationURL != test.mockResponse.Previous {
					t.Errorf("Expected previousLocationURL to be updated")
				}
			}

		})
	}
}

func TestCommandMapPagination(t *testing.T) {
	nextURL := "https://pokeapi.co/api/v2/location-area?offset=20"
	prevURL := "https://pokeapi.co/api/v2/location-area?offset=0"

	cfg := &config{
		pokeClient: &mockClient{
			getLocationAreasFunc: func(pageURL *string) (pokeapi.LocationAreaResponse, error) {
				return pokeapi.LocationAreaResponse{
					Next:     &nextURL,
					Previous: &prevURL,
				}, nil
			},
		},
	}

	err := commandMap(cfg, "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if cfg.nextLocationURL == nil || *cfg.nextLocationURL != nextURL {
		t.Errorf("Expected nextLocationURL %s, got %v", nextURL, cfg.nextLocationURL)
	}
	if cfg.previousLocationURL == nil || *cfg.previousLocationURL != prevURL {
		t.Errorf("Expected previsousLocationURL %s, got %v", prevURL, cfg.previousLocationURL)
	}

}
