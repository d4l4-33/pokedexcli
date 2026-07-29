package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonGet(pokemonName string) (PokemonStruct, error) {
	url := baseURL + "/pokemon/" + pokemonName + "/"

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonStruct{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonStruct{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStruct{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStruct{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStruct{}, err
	}

	pokemonResp := PokemonStruct{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonStruct{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
