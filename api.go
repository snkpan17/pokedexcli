package main

import (
	"encoding/json"
	"fmt"
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
