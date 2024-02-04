package main

import (
	"fmt"

	"github.com/Woutjeee/go_pokedex/internal/commands"
)

func main() {
	fmt.Println("Hello, World!")

	commands := commands.GetCommands()

	fmt.Println(commands)
}
