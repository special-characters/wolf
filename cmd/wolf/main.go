package main

import (
	"fmt"
	"os"

	"github.com/special-characters/wolf"
)

func main() {
	lexicon := wolf.New()

	if len(os.Args) > 1 {
		if wolf.IsValidHex(os.Args[1]) {

			// since arg has passed validation, we don't anticipate any errors
			name, _ := lexicon.Wolf(os.Args[1])

			fmt.Println(name)
		} else {

			hex, err := lexicon.Unwolf(os.Args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println(hex)
		}
	} else {
		fmt.Println(lexicon.Random(2))
	}
}
