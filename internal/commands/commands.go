package commands

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func CommandHelp() error {
	fmt.Println("Help commands")
	return nil
}

func CommandExit() error {
	fmt.Println("Help commands")
	return nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}
}
