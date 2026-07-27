package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var locationArea config = config{
	Next: "https://pokeapi.co/api/v2/location-area/",
}

func commandMap(c *config) error {
	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder((res.Body))
	if err := decoder.Decode(&locationArea); err != nil {
		return err
	}

	for _, location := range locationArea.Results {
		fmt.Println(location)
	}

	return nil
}
