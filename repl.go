package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Printf("error creating new scanner: %s", scanner.Err())
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command, exists := getCommands()[strings.ToLower(input)]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower((text))
	return strings.Fields(lowerCase)
}

/*  #####ARCHIVE######

startRepl()
scanner := bufio.NewScanner(os.Stdin)
if scanner.Err() != nil {
	fmt.Println("error creating scanner")
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		if len(text) == 0 {
			continue
			}

			text_clean := cleanInput(text)

			fmt.Printf("Your command was: %s\n", text_clean[0])
			}


*/
