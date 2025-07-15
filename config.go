package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Next     string
	Previous string
}

var config Config

//var config = Config{
//	Next:     os.Getenv("POKE_LOCATION_URL"),
//	Previous: "",
//}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("error loading .env file: %v", err)
		os.Exit(1)
	}
	config = Config{
		Next:     os.Getenv("POKE_LOCATION_URL"),
		Previous: "",
	}
	//fmt.Printf("from env: %s\n", os.Getenv("POKE_LOCATION_URL"))

}
