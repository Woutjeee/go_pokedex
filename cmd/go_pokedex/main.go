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
	fmt.Println("Hello, World!")

	commands := commands.GetCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			fmt.Println(command)
		} else {
			fmt.Println("Does not exist")
		}
		printPrompt()
	}
}
