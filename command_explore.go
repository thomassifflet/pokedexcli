package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location name provided")
	}
	locationName := args[0]

	location, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", location.Name)

	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
