package main

import (
	"fmt"
	"github.com/snkpan17/pokedexcli/internal/pokecache"
	"os"
	"time"
)

func printAllDesc() {
	for _, cmd := range commands {
		fmt.Println(cmd.name + ": " + cmd.desc)
	}
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*Config, []string) error
}

var commands map[string]cliCommand
var cache *pokecache.Cache

func commandExit(conf *Config, words []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	cache.Stop()
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config, words []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	printAllDesc()
	return nil
}

func commandMap(conf *Config, words []string) error {
	url := conf.Next
	locations, prev, next, err := getLocations(url, cache)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return err
	}
	for _, loc := range locations {
		fmt.Println(loc)
	}
	conf.Previous = prev
	conf.Next = next
	return nil
}

func commandMapB(conf *Config, words []string) error {
	url := conf.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		conf.Previous = ""
		conf.Next = os.Getenv("POKE_LOCATION_URL")
		return nil
	}
	locations, prev, next, err := getLocations(url, cache)
	if err != nil {
		return err
	}
	for _, loc := range locations {
		fmt.Println(loc)
	}
	conf.Previous = prev
	conf.Next = next
	return nil
}

func commandExplore(c *Config, words []string) error {
	loc := words[1]
	url := c.POKE_LOCATION_BASE_URL + loc
	fmt.Printf("Exploring %s...\n", loc)
	pokemons, err := getPokeAtLoc(url, cache)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Println(" - " + pokemon)
	}
	return nil
}

func commandCatch(c *Config, words []string) error {
	pokemon := words[1]
	url := c.POKEMON_BASE_URL + pokemon
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	Pokemon, err := getPokeDetail(url, cache)
	pokeExp := Pokemon.Base_experience
	if err != nil {
		return err
	}
	userDefeatsPoke := canDefeat(c.UserExp, pokeExp)
	if userDefeatsPoke {
		fmt.Println(pokemon + " was caught!")
		c.Pokedex[pokemon] = Pokemon
		if pokeExp > c.UserExp {
			c.UserExp = pokeExp
		}
	} else {
		fmt.Println(pokemon + " escaped!")
	}
	return nil
}

func printPokemon(Pokemon PokemonDetailResponse) {
	fmt.Println("Name: ", Pokemon.Name)
	fmt.Println("Height: ", Pokemon.Height)
	fmt.Println("Weight: ", Pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range Pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range Pokemon.Types {
		fmt.Println(" - ", t.Type.Name)
	}
}

func commandInspect(c *Config, words []string) error {
	pokemon := words[1]
	Pokemon, ok := c.Pokedex[pokemon]
	if !ok {
		fmt.Println("Pokemon not found in Pokedex. Not caught")
	} else {
		printPokemon(Pokemon)
	}
	return nil
}

func commandPokedex(c *Config, words []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range c.Pokedex {
		fmt.Println(" - ", key)
	}
	return nil
}

func init() {
	cache = pokecache.NewCache(5 * time.Second)
	commands = make(map[string]cliCommand)
	commands["help"] = cliCommand{
		name:     "help",
		desc:     "Displays a help message",
		callback: commandHelp,
	}
	commands["exit"] = cliCommand{
		name:     "exit",
		desc:     "Exit the Pokedex",
		callback: commandExit,
	}
	commands["map"] = cliCommand{
		name:     "map",
		desc:     "Show 20 locations in the Pokemon world",
		callback: commandMap,
	}
	commands["mapb"] = cliCommand{
		name:     "mapb",
		desc:     "Show previous 20 locations in the Pokemon world",
		callback: commandMapB,
	}
	commands["explore"] = cliCommand{
		name:     "explore",
		desc:     "Show pokemons at this area",
		callback: commandExplore,
	}
	commands["catch"] = cliCommand{
		name:     "catch",
		desc:     "Throw pokeball to catch a pokemon",
		callback: commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:     "inspect",
		desc:     "Get caught pokemon details",
		callback: commandInspect,
	}
	commands["pokedex"] = cliCommand{
		name:     "pokedex",
		desc:     "Show pokemons caught",
		callback: commandPokedex,
	}
}
