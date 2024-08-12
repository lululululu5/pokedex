package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Name string `json:"name"`
}

type Pokedex struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Result `json:"results"`
}

// Split up the model from the request. Build the Pokemon response model and then create the request for it. 
// move the extraction functionality into the request

var pokedex Pokedex

func commandMap(pokedex *Pokedex) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	
	err = json.Unmarshal(body, pokedex)
	if err != nil {
		log.Fatal((err))
	}

	for _, result := range pokedex.Results {
		fmt.Println(result.Name)
	}
}

func commandBMap() error {
	fmt.Println("Show previous 20 locations")
	return nil
}