package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationAreaResponse LocationAreaResponse

	resp, err := c.httpClient.Get(url)
	if err != nil {
		err := fmt.Errorf("Error making GET request to %s: %w", url, err)
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Received a non-OK HTTP status code: %d", resp.StatusCode)
		return LocationAreaResponse{}, err
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreaResponse); err != nil {
		err := fmt.Errorf("Error decoding JSON response: %w", err)
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil

}
