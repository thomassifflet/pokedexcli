package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s ...\n", pokemonName)
	time.Sleep(50 * time.Millisecond)
	const threshold = 50
	randNumber := rand.Intn(pokemon.BaseExperience)
	if randNumber > threshold {
		fmt.Printf("%s escaped !\n", pokemonName)
	}

	cfg.pokemonCaught[pokemonName] = pokemon
	fmt.Printf("%s was captured !\n", pokemonName)

	return nil
}
