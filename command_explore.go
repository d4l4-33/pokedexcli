package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	area := args[0]
	locationResp, err := cfg.pokeapiClient.ListPokemon(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, loc := range locationResp.PokemonEncounters {
		fmt.Printf("- %s\n", loc.Pokemon.Name)
	}
	return nil
}
