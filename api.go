package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snkpan17/pokedexcli/internal/pokecache"
)

type location struct {
	Name string
	Url  string
}

type LocationApiResponse struct {
	Count    uint       `json:"count"`
	Previous string     `json:"previous"`
	Next     string     `json:"next"`
	Results  []location `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
	Version any     `json:"version_details"`
}

type PokemonResponse struct {
	Pokemon_encounters []PokemonEncounters `json:"pokemon_encounters"`
}

func getLocations(url string, cache *pokecache.Cache) ([]string, string, string, error) {
	var Data LocationApiResponse
	var body []byte
	cached, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, "", "", err
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []string{}, "", "", err
		}
		cache.Add(url, body)
	} else {
		body = cached
	}
	err := json.Unmarshal(body, &Data)
	if err != nil {
		return []string{}, "", "", err
	}
	var locations []string
	for _, loc := range Data.Results {
		locations = append(locations, loc.Name)
	}
	return locations, Data.Previous, Data.Next, nil

}

func getPokeAtLoc(url string, cache *pokecache.Cache) ([]string, error) {
	var Data PokemonResponse
	var body []byte
	cached, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, err
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []string{}, err
		}
		cache.Add(url, body)
	} else {
		body = cached
	}
	err := json.Unmarshal(body, &Data)
	if err != nil {
		return []string{}, err
	}
	var pokemons []string
	for _, data := range Data.Pokemon_encounters {
		pokemons = append(pokemons, data.Pokemon.Name)
	}
	return pokemons, nil
}
