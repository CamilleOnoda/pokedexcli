package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemonInLocationArea(pageURL *string) (PokemonInLocationResponse, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var svc Service = &ServiceImpl{client: *c}
	data, err := svc.GetData(url)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}

	var pokemonInLocationResponse PokemonInLocationResponse

	if err := json.Unmarshal(data, &pokemonInLocationResponse); err != nil {
		err := fmt.Errorf("Error unmarshaling JSON response: %w", err)
		return PokemonInLocationResponse{}, err
	}

	return pokemonInLocationResponse, nil
}
