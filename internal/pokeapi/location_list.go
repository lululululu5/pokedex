package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(url)
	if ok {
		// casche hit
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
	
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}

func (c *Client) ListLocationPokemon(userInputLocation string) (RespLocationPokemons, error) {
	url := baseURL + "/location-area/" + userInputLocation

	// check caching if it exists
	dat, ok := c.cache.Get(url)
	if ok {
		//cache hit
		pokemonResp := RespLocationPokemons{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return RespLocationPokemons{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	pokemonResp := RespLocationPokemons{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	// add chaching
	c.cache.Add(url, dat)

	return pokemonResp ,nil
}