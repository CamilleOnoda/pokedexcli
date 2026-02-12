package pokeapi

import (
	"fmt"
	"net/http"
	"time"

	pokecache "github.com/CamilleOnoda/pokedexcli/internal/pokecache"
)

type PokeAPIClient interface {
	GetLocationAreas(pageURL *string) (LocationAreaResponse, error)
	GetPokemonInLocationArea(pageURL *string) (PokemonInLocationResponse, error)
	GetPokemonNamesInLocationArea(areaName string) (PokemonInLocationResponse, error)
	GetPokemonInfo(pokemonName string) (Pokemon, error)
}

type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(10 * time.Minute),
	}
}

func (c *Client) GetPokemonNamesInLocationArea(areaName string) (PokemonInLocationResponse, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := fmt.Sprintf("%s/location-area/%s/", baseURL, areaName)

	return c.GetPokemonInLocationArea(&url)
}
