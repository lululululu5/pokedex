package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lululululu5/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient		pokeapi.Client
	pokedex				pokeapi.Pokedex
	nextLocationsURL	*string
	prevLocationsURL	*string
}


func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		input := ""
		if len(words) == 2 {
			input = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, input)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words

}

type cliCommand struct {
	name string
	description string
	callback func(conf *config, input string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help", 
			description: "Display a help message",
			callback: commandHelp,
		},

		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},

		"map": {
			name: "map",
			description: "Show next 20 locations",
			callback: commandMapf,
		},
		
		"mapb": {
			name: "mapb", 
			description: "Show previous 20 locations",
			callback: commandMapb,
		},

		"explore": {
			name: "explore", 
			description: `Explore area for specific pokemeon`,
			callback: commandExplore,
		},

		"catch": {
			name: "catch", 
			description: `Catch pokemon`,
			callback: commandCatch,
		},

		"inspect": {
			name: "inspect",
			description: "See stats for Pokemon",
			callback: commandInspect,
		},

		"pokedex": {
			name: "pokedex",
			description: "See list of caught Pokemons ",
			callback: commandPokedex,
		},
	}
}