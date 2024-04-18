package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("You have caught these Pokemon: \n")
	for _, p := range cfg.pokemonCaught {

		fmt.Printf("- %s\n", p.Name)

	}

	return nil
}
