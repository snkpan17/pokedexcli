package main

import (
	"fmt"
	"os"
)

func printAllDesc() {
	for _, cmd := range commands {
		fmt.Println(cmd.name + ": " + cmd.desc)
	}
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*Config) error
}

var commands map[string]cliCommand

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	printAllDesc()
	return nil
}

func commandMap(conf *Config) error {
	url := conf.Next
	locations, prev, next, err := getLocations(url)
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

func commandMapB(conf *Config) error {
	url := conf.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		conf.Previous = ""
		conf.Next = os.Getenv("POKE_LOCATION_URL")
		return nil
	}
	locations, prev, next, err := getLocations(url)
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

func init() {
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
}
