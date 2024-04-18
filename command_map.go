package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas :")

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationUrl = resp.Previous

	return nil
}

func callbackMapBack(cfg *config) error {
	if cfg.previousLocationUrl == nil {
		err := errors.New("you are on the first page, sorry")
		return err
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas :")

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationUrl = resp.Previous

	return nil
}
