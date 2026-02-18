package pokeapi

import (
	"time"
)

// Wraps a PokeAPIClient and add caching
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
		key = "locations" + *pageURL
	}

	if cached, found := c.cache.Get(key); found {
		return cached.(LocationAreaResponse), nil
	}

	resp, err := c.client.GetLocationAreas(pageURL)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Set(key, resp, c.ttl)

	return resp, nil

}

func (c *CachedClient) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	key := "pokemon:" + pokemonName

	if cached, found := c.cache.Get(key); found {
		return cached.(Pokemon), nil
	}

	resp, err := c.client.GetPokemonInfo(pokemonName)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Set(key, resp, c.ttl)

	return resp, nil

}

func (c *CachedClient) GetPokemonInLocationArea(areaURL *string) (PokemonInLocationResponse, error) {
	if areaURL == nil {
		return PokemonInLocationResponse{}, nil
	}

	key := "location:" + *areaURL

	if cached, found := c.cache.Get(key); found {
		return cached.(PokemonInLocationResponse), nil
	}

	resp, err := c.client.GetPokemonInLocationArea(areaURL)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}

	c.cache.Set(key, resp, c.ttl)

	return resp, nil
}
