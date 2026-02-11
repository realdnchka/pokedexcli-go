package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/realdnchka/pokedexcli-go/internal/pokectx"
)

var dbPokemon pokectx.DB
var ctx = pokectx.Context{
	Database: &dbPokemon,
}

func main() {
	initCommands()
	dbPokemon.Create()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)

		if len(input) != 0 {
			if c, ok := supportedCommands[input[0]]; ok {
				c.callback(&ctx, input[1:])
				continue
			} else {
				log.Printf("No such command")
				continue
			}
		}
	}
}