package main

import (
	"encoding/json"
	"io"
	"net/http"
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

func getLocations(url string) ([]string, string, string, error) {
	res, err := http.Get(url)
	if err != nil {
		return []string{}, "", "", err
	}
	var Data LocationApiResponse
	decoder := json.NewDecoder(res.Body)
	for {
		if err := decoder.Decode(&Data); err == io.EOF {
			break
		} else if err != nil {
			return []string{}, "", "", err
		}
	}
	var locations []string
	for _, loc := range Data.Results {
		locations = append(locations, loc.Name)
	}
	return locations, Data.Previous, Data.Next, nil

}
