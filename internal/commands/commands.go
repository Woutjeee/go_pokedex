package commands

import (
	"github.com/Woutjeee/go_pokedex/internal/commands/exit"
	"github.com/Woutjeee/go_pokedex/internal/commands/help"
	"github.com/Woutjeee/go_pokedex/internal/commands/types"
)

func GetCommands() map[string]types.CliCommand {
	return map[string]types.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    help.CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    exit.CommandExit,
		},
	}
}
