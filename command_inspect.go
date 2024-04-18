package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.pokemonCaught[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Type(s):")
	for _, ptype := range pokemon.Types {
		fmt.Printf(" - %s\n", ptype.Type.Name)
	}

	return nil
}
