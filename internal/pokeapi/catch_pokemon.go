package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokedex struct{
	Pokemons map[string]Pokemon
}

type Pokemon struct{
	Name string
	BaseExperience int

}

type RespPokemon struct{
	BaseExperience int `json:"base_experience"`
	Name string `json:"name"`
}

func (c *Client) GetPokemonInfo(userInputPokemon string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + userInputPokemon

	// add caching
	dat, ok := c.cache.Get(url)
	if ok {
		println("EXISTING CACHE")
		pokemonData := RespPokemon{}
		err := json.Unmarshal(dat, &pokemonData)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokemonData, err
	}
	println("NO CACHING")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonData := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokemonData, err
}