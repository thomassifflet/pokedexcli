package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Pokedex help menu !")
	fmt.Println("Commands available :")

	commandList := getCommands()
	for _, cmd := range commandList {
		fmt.Printf(" - %s: %s", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
