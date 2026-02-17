package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/pokemon/" + pokemonName + "/"

	var svc Service = &ServiceImpl{client: *c}
	data, err := svc.GetData(url)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		err := fmt.Errorf("Error unmarshaling JSON response: %w", err)
		return Pokemon{}, err
	}

	return pokemon, nil
}
