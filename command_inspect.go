package main

import (
	"errors"
	"fmt"
)

// Update RespPokemon and data held in pokedex - Check
// Create function to get data from Pokedex in case value exists
// Add function to repl - Check

func commandInspect(cfg *config, input string) error {
	if input == "" {
		return errors.New("please provide one pokemon to inspect")
	}

	pokemon, ok:= cfg.pokedex.Pokemons[input]
	if !ok {
		return errors.New("pokemon not in pokedex, you need to catch it first")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for k,v := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", k, v)
	}
	fmt.Println("Types:")
	for _,val := range pokemon.Types {
		fmt.Printf(" - %s\n", val)
	}
	
	return nil
}

func commandPokedex(cfg *config, input string) error {
	pokemons := cfg.pokedex.Pokemons

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemons {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}