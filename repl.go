package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Print help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Print a list of the possible locations",
			callback:    callbackMap,
		},
		"mapback": {
			name:        "mapback",
			description: "Print the list of previously displayed locations",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {location_name}",
			description: "Give a list of the Pokemon in that area",
			callback:    callbackExplore,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "List the details of caught {pokemon_name}",
			callback:    callbackInspect,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "try to catch a pokemon with name {pokemon_name}",
			callback:    callbackCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(" >")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Command not valid")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(str string) []string {
	loweredInput := strings.ToLower(str)
	words := strings.Fields(loweredInput)
	return words
}
