package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Woutjeee/go_pokedex/internal/commands"
)

var cliName string = "Pokedex"

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	commands := commands.GetCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			err := command.Callback()

			if err != nil {
				fmt.Println("Error", err)
			}

			// Break the loop if the command is exit.
			if command.Name == "exit" {
				break
			}

		} else {
			fmt.Println("Does not exist")
		}
		printPrompt()
	}
}
