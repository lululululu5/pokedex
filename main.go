package main

import (
	"time"

	"github.com/lululululu5/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	
	startRepl(&cfg)
}