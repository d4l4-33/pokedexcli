package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationInfo struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Next     string
	Previous string
}

var locationAreaURL config = config{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}

func commandMap(c *config) error {
	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var location locationInfo
	decoder := json.NewDecoder((res.Body))
	if err := decoder.Decode(&location); err != nil {
		return err
	}

	c.Next = location.Next
	c.Previous = location.Previous

	for _, location := range location.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(c.Previous)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var location locationInfo
	decoder := json.NewDecoder((res.Body))
	if err := decoder.Decode(&location); err != nil {
		return err
	}

	c.Next = location.Next
	c.Previous = location.Previous

	for _, location := range location.Results {
		fmt.Println(location.Name)
	}

	return nil
}
