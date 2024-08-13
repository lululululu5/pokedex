package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, input string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL) // returns RespShallowLocations struct
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, input string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results{
		fmt.Println(loc.Name)
	}
	return nil
}