package main

import (
	"time"

	"github.com/lululululu5/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	
	startRepl(cfg)
}