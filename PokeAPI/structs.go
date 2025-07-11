package pokeapi

//all the structs that will be used for json data

type locations struct {
	Next    *string    `json:"next"`
	Prev    *string    `json:"previous"`
	Results []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type area struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
