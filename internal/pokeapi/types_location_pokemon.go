package pokeapi

type PokemonInLocationResponse struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	GameIndex         int    `json:"game_index"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
