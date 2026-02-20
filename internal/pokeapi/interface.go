package pokeapi

// Define contract for interacting with PokeAPI
// Both Client and CachedClient will implement

type PokeAPIClient interface {
	GetLocationAreas(pageURL *string) (LocationAreaResponse, error)
	GetPokemonInfo(pokemonName string) (Pokemon, error)
	GetPokemonInLocationArea(areaName *string) (PokemonInLocationResponse, error)
	Clear()
}
