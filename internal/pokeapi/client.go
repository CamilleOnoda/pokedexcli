package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "github.com/CamilleOnoda/pokedexcli/internal/pokecache"
)

type PokeAPIClient interface {
	GetLocationAreas(pageURL *string) (LocationAreaResponse, error)
	GetPokemonInLocationArea(pageURL *string) (PokemonInLocationResponse, error)
	BuildURLPokemonInLocationArea(areaName string) (*string, error)
	GetPokemonInfo(pokemonName string) (Pokemon, error)
}

type Client struct {
	httpClient *http.Client
	Cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(10 * time.Minute),
	}
}

func (c *Client) FetchData(url string) ([]byte, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		err := fmt.Errorf("Error making GET request to %s: %w", url, err)
		return []byte{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Received a non-OK HTTP status code: %d", resp.StatusCode)
		return []byte{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Error reading response body: %w", err)
		return []byte{}, err
	}

	return body, nil
}
