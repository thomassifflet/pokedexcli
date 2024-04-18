package main

import "github.com/thomassifflet/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	previousLocationUrl *string
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&config)

}
