package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("you need to input a pokemon name")
		return nil
	}

	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Println("pokemon not in pokedex")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pkmType := range pokemon.Types {
		fmt.Printf("  - %s\n", pkmType.Type.Name)
	}

	return nil
}
