package pokeapi

// Wraps any client implementation and add caching

import (
	"encoding/json"
	"time"
)

type CachedClient struct {
	client PokeAPIClient
	cache  Cache
	ttl    time.Duration
}

func NewCachedClient(client PokeAPIClient, cache Cache, ttl time.Duration) *CachedClient {
	return &CachedClient{
		client: client,
		cache:  cache,
		ttl:    ttl,
	}
}

func (c *CachedClient) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	key := "locations:default"
	if pageURL != nil {
		key = "locations:" + *pageURL
	}

	if cached, found := c.cache.Get(key); found {
		var resp LocationAreaResponse
		if err := json.Unmarshal(cached, &resp); err == nil {
			return resp, nil
		}
	}

	resp, err := c.client.GetLocationAreas(pageURL)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if data, err := json.Marshal(resp); err == nil {
		c.cache.Set(key, data, c.ttl)
	}

	return resp, nil

}

func (c *CachedClient) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	key := "pokemon:" + pokemonName

	if cached, found := c.cache.Get(key); found {
		var resp Pokemon
		if err := json.Unmarshal(cached, &resp); err == nil {
			return resp, nil
		}
	}

	resp, err := c.client.GetPokemonInfo(pokemonName)
	if err != nil {
		return Pokemon{}, err
	}

	if data, err := json.Marshal(resp); err == nil {
		c.cache.Set(key, data, c.ttl)
	}

	return resp, nil

}

func (c *CachedClient) GetPokemonInLocationArea(areaURL *string) (PokemonInLocationResponse, error) {
	if areaURL == nil {
		return PokemonInLocationResponse{}, nil
	}

	key := "location:" + *areaURL

	if cached, found := c.cache.Get(key); found {
		var resp PokemonInLocationResponse
		if err := json.Unmarshal(cached, &resp); err == nil {
			return resp, nil
		}
	}

	resp, err := c.client.GetPokemonInLocationArea(areaURL)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}
	if data, err := json.Marshal(resp); err == nil {
		c.cache.Set(key, data, c.ttl)
	}

	return resp, nil
}

func (c *CachedClient) Clear() {
	c.cache.Clear()
}
