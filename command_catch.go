package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/lululululu5/pokedexcli/internal/pokeapi"
)


func commandCatch(cfg *config, input string) error {
	if input == "" {
		return errors.New("please provide a single pokemon")
	}

	userInputPokemon := input
	RespPokemon, err := cfg.pokeapiClient.GetPokemonInfo(userInputPokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", userInputPokemon)

	pokemonExperience := RespPokemon.BaseExperience
	userLevel := 100

	if userLevel > rand.IntN(pokemonExperience) {
		fmt.Printf("%s was caught!\n", userInputPokemon)
		if _, ok := cfg.pokedex.Pokemons[RespPokemon.Name]; !ok {
			stats := make(map[string]int)
			for _,val := range RespPokemon.Stats {
				stats[val.Stat.Name] = val.BaseStat 
			}
			types := []string{}
			for _,val := range RespPokemon.Types {
				types = append(types, val.Type.Name)

			}
			cfg.pokedex.Pokemons[RespPokemon.Name] = pokeapi.Pokemon{
				Name: RespPokemon.Name,
				BaseExperience: RespPokemon.BaseExperience,
				Weight: RespPokemon.Weight,
				Height: RespPokemon.Height,
				Stats: stats,
				Types: types,
			}
		}
	} else {
		fmt.Printf("%s escaped!\n", userInputPokemon)
	}

	return nil

}