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
			cfg.pokedex.Pokemons[RespPokemon.Name] = pokeapi.Pokemon{
				Name: RespPokemon.Name,
				BaseExperience: RespPokemon.BaseExperience,
			}
		}
		fmt.Print(cfg.pokedex)
	} else {
		fmt.Printf("%s escaped!\n", userInputPokemon)
	}

	return nil

}