package main

import (
	"time"

	"github.com/thomassifflet/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	previousLocationUrl *string
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&config)

}
