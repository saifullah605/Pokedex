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

type pokemon struct {
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
