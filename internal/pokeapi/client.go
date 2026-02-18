package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		baseURL:    "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error making GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("Received status: %w", err)
	}

	var locationResp LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationResp); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error decoding JSON %w", err)
	}

	return locationResp, nil

}

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := c.baseURL + "/pokemon/" + pokemonName + "/"

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error fetching Pokemon: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("Received status: %w", err)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("Error decoding JSON: %w", err)
	}

	return pokemon, nil

}

func (c *Client) GetPokemonInLocationArea(areaName *string) (PokemonInLocationResponse, error) {
	url := fmt.Sprintf("%s/location-area/%v/", c.baseURL, areaName)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonInLocationResponse{}, fmt.Errorf("Error fetching Pokemon: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PokemonInLocationResponse{}, fmt.Errorf("Received status: %w", err)
	}

	var pokemonInLocationResp PokemonInLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokemonInLocationResp); err != nil {
		return PokemonInLocationResponse{}, fmt.Errorf("Error decoding JSON: %w", err)
	}

	return pokemonInLocationResp, nil

}
