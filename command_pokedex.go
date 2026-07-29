package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for pokeName, _ := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokeName)
	}
	return nil
}
