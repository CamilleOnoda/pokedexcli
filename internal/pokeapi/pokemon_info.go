package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/pokemon/" + pokemonName + "/"

	var pokemon Pokemon

	resp, err := c.httpClient.Get(url)
	if err != nil {
		err := fmt.Errorf("Error making GET request to %s: %w", url, err)
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Received a non-OK HTTP status code: %d", resp.StatusCode)
		return Pokemon{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Error reading response body: %w", err)
		return Pokemon{}, err
	}

	if err := json.Unmarshal(body, &pokemon); err != nil {
		err := fmt.Errorf("Error unmarshaling JSON response: %w", err)
		return Pokemon{}, err
	}

	return pokemon, nil
}
