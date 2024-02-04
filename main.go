package main

import (
	"fmt"
)

var cliName string = "Pokedex"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Help")
	return nil
}

func commandExit() error {
	fmt.Println("Help")
	return nil
}

func printPrompt() {
	fmt.Println(cliName, "> ")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
}
