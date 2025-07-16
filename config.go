package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Next                   string
	Previous               string
	POKE_LOCATION_BASE_URL string
}

var config Config

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("error loading .env file: %v", err)
		os.Exit(1)
	}
	config = Config{
		Next:                   os.Getenv("POKE_LOCATION_URL"),
		Previous:               "",
		POKE_LOCATION_BASE_URL: os.Getenv("POKE_LOCATION_BASE_URL"),
	}

}
