package pokeapi

type Pokedex struct{
	Pokemons map[string]Pokemon
}

type Pokemon struct{
	Name string
	BaseExperience int
	Weight int
	Height int
	Stats map[string]int
	Types []string


}


type RespPokemon struct {
	BaseExperience int `json:"base_experience"`
	Name string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
	Height int `json:"height"`
}