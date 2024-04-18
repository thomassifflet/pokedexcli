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
	callback    func(*config) error
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
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Command not valid")
			continue
		}

		err := command.callback(cfg)
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
