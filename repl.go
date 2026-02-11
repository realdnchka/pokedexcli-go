package main

import (
	"fmt"
	"log"
	"math/rand"
	"math"
	"os"
	"strings"

	"github.com/realdnchka/pokedexcli/internal/pokeapi"
	"github.com/realdnchka/pokedexcli/internal/pokectx"
)

type Command struct {
	name string
	description string
	callback func(ctx *pokectx.Context, args []string)
}

var supportedCommands map[string]Command

func initCommands() {
	supportedCommands = map[string]Command {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Shows a list of areas",
			callback: commandMapNext,
		},
		"mapb": {
			name: "mapb",
			description: "Shows a list of areas",
			callback: commandMapPrev,
		},
		"explore": {
			name: "explore",
			description: "Shows list of Pokemon at the <area> param",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Try to catch a Pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect yours Pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all of your pokemons",
			callback: commandPokedex,
		},
	}
}

func commandHelp(ctx *pokectx.Context, args []string) {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	sc := supportedCommands
	for _, c := range sc {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
}

func commandExit(ctx *pokectx.Context, args []string) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandMapNext(ctx *pokectx.Context, args []string) {
	areas, err := pokeapi.GetAreas(true)
	if err != nil {
		log.Printf("cannot get response: %v\n", err)
		return
	}
	for _, a := range areas.Areas {
		fmt.Printf("%v\n", a.Name)
	}
}

func commandMapPrev(ctx *pokectx.Context, args []string) {
	areas, err := pokeapi.GetAreas(false)
	if err != nil {
		log.Printf("cannot get response: %v\n", err)
	}
	for _, a := range areas.Areas {
		fmt.Printf("%v\n", a.Name)
	}
}

func commandExplore(ctx *pokectx.Context, area []string) {
	poks, err := pokeapi.GetAreaPokemons(area[0])
	if err != nil {
		log.Printf("cannot get response: %v\n", err)
		return
	}

	for _, p := range poks.PokemonEncounters {
		fmt.Printf("%v\n", p.Pokemon.Name)
	}
}

func commandCatch(ctx *pokectx.Context, args []string) {
	var stats []pokectx.PokemonStat
	var types []pokectx.PokemonType

	pokemon := args[0]
	
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	
	p, err := pokeapi.GetPokemon(pokemon)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	for _, s := range p.Stats {
		stats = append(stats, pokectx.PokemonStat { 
			Name: s.Name,
			Value: s.Value,
		})
	}

	for _, t := range p.Types {
		types = append(types, pokectx.PokemonType{
			Name: t.Name,
		})
	}
	pok := pokectx.Pokemon {
		ID: p.ID,
		Name: p.Name,
		BaseEXP: p.BaseEXP,
		Height: p.Height,
		Weight: p.Weight,
		Stats: stats,
		Types: types,
	}

	if catch(pok.BaseEXP) {
		fmt.Printf("%s was caught!\n", pok.Name)
		ctx.Database.Write(p.Name, pok)
	} else {
		fmt.Printf("%s escaped!\n", pok.Name)
	}
}

func commandInspect(ctx *pokectx.Context, args []string) {
	p := pokectx.Pokemon {
		Name: args[0],
	}

	p, ok := ctx.Database.Read(p.Name); 
	if !ok {
		fmt.Printf("You dont have such Pokemon\n")
		return
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %v\n", p.Height)
	fmt.Printf("Weight: %v\n", p.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range p.Stats{
		fmt.Printf("\t-%s: %v\n", s.Name, s.Value)
	}
	fmt.Printf("Types:\n")
	for _, s := range p.Types{
		fmt.Printf("\t- %s\n", s.Name)
	}
	
}

func commandPokedex(ctx *pokectx.Context, args []string) {
	if len(ctx.Database.ReadAll()) == 0 {
		fmt.Printf("You didn't catch any Pokemons yet\n")
		return
	}
	fmt.Printf("Your Pokedex:\n")
	for _, p := range ctx.Database.ReadAll() {
		fmt.Printf("\t- %s\n", p.Name)
	}
}

func catch(d int) bool {
	chance := math.Exp(-float64(d)/300)
	return rand.Float64() < chance
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(input)))
}