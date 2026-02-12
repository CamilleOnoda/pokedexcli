package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetPokemonInLocationArea(pageURL *string) (PokemonInLocationResponse, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheStart := time.Now()
	cacheKey := url
	if cachedData, found := c.cache.Get(cacheKey); found {
		elapsed := time.Since(cacheStart)
		fmt.Printf("CACHE HIT: retrieved data from cache in %v\n", elapsed)
		var cachedResponse PokemonInLocationResponse
		if err := json.Unmarshal(cachedData, &cachedResponse); err == nil {
			return cachedResponse, nil
		}
	}

	var pokemonInLocationResponse PokemonInLocationResponse

	reqStart := time.Now()
	resp, err := c.httpClient.Get(url)
	if err != nil {
		err := fmt.Errorf("Error making GET request to %s: %w", url, err)
		return PokemonInLocationResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Received a non-OK HTTP status code: %d", resp.StatusCode)
		return PokemonInLocationResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Received a non-OK HTTP status code: %d", resp.StatusCode)
		return PokemonInLocationResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Error reading response body: %w", err)
		return PokemonInLocationResponse{}, err
	}

	c.cache.Add(cacheKey, body)

	if err := json.Unmarshal(body, &pokemonInLocationResponse); err != nil {
		err := fmt.Errorf("Error unmarshaling JSON response: %w", err)
		return PokemonInLocationResponse{}, err
	}

	elapsed := time.Since(reqStart)
	fmt.Printf("API CALL: fetched data from API in %v\n", elapsed)

	return pokemonInLocationResponse, nil
}
