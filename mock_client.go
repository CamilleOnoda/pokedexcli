package main

import (
	pokeapi "github.com/CamilleOnoda/pokedexcli/internal/pokeapi"
)

// mock client to test commands without hitting the real API

type mockClient struct {
	getLocationAreasFunc         func(pageURL *string) (pokeapi.LocationAreaResponse, error)
	getPokemonInfoFunc           func(pokemonName string) (pokeapi.Pokemon, error)
	getPokemonInLocationAreaFunc func(areaURL *string) (pokeapi.PokemonInLocationResponse, error)
}

func (m *mockClient) GetLocationAreas(pageURL *string) (pokeapi.LocationAreaResponse, error) {
	if m.getLocationAreasFunc == nil {
		return pokeapi.LocationAreaResponse{}, nil
	}
	return m.getLocationAreasFunc(pageURL)
}

func (m *mockClient) GetPokemonInfo(pokemonName string) (pokeapi.Pokemon, error) {
	if m.getPokemonInfoFunc == nil {
		return pokeapi.Pokemon{}, nil
	}
	return m.getPokemonInfoFunc(pokemonName)
}

func (m *mockClient) GetPokemonInLocationArea(areaURL *string) (pokeapi.PokemonInLocationResponse, error) {
	if m.getPokemonInLocationAreaFunc == nil {
		return pokeapi.PokemonInLocationResponse{}, nil
	}
	return m.getPokemonInLocationAreaFunc(areaURL)
}

func (m *mockClient) Clear() {}
