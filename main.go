package main

import (
	"time"

	"github.com/thomassifflet/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	previousLocationUrl *string
	pokemonCaught       map[string]pokeapi.Pokemon
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokemonCaught: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&config)

}
