package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/CamilleOnoda/pokedexcli/internal/pokecache"
)

type PokeAPIClient interface {
	GetLocationAreas(pageURL *string) (LocationAreaResponse, error)
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
