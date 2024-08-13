package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, input string) error {
	if input == "" {
		return errors.New("please provide a single location name")
	}
	userInputLocation := input
	RespLocationPokemons, err := cfg.pokeapiClient.ListLocationPokemon(userInputLocation) // returns RespLocationPokemons
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", userInputLocation)
	fmt.Println("Found Pokemon: ")
	
	for _, encounter := range RespLocationPokemons.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}