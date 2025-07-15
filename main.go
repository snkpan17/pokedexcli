package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("error loading .env file: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			command, ok := commands[words[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			if err := command.callback(&config); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func cleanInput(input string) []string {
	str := strings.ToLower(input)
	str = strings.Trim(str, " ")
	output := strings.Split(str, " ")
	var filtered []string
	for _, s := range output {
		if s == "" {
			continue
		}
		filtered = append(filtered, s)
	}
	return filtered
}
