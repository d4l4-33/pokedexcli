package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("pokemon name required")
		return nil
	}

	inputName := args[0]
	pokemon, err := cfg.pokeapiClient.PokemonGet(inputName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Intn(350) < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.pokedex[pokemon.Name] = pokemon
	_, ok := cfg.pokedex[pokemon.Name]
	if ok {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	}
	return errors.New("error: pokedex failed to register pokemon")
}
