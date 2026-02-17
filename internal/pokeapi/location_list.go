package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var svc Service = &ServiceImpl{client: *c}
	data, err := svc.GetData(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationAreaResponse LocationAreaResponse
	if err := json.Unmarshal(data, &locationAreaResponse); err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResponse, nil

}
