package pokeapi

import "fmt"

func (c *Client) BuildURLPokemonInLocationArea(areaName string) (*string, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := fmt.Sprintf("%s/location-area/%s/", baseURL, areaName)
	return &url, nil
}
