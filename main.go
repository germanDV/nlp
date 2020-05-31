// Generates Natural Language Passwords based on "diceware"
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gitlab.com/germanDV/nlp/csvreader"
	"gitlab.com/germanDV/nlp/shaker"
)

// findWord finds a word in a CSV file by its ID
func findWord(n int, dice int) string {
	var filename string
	if dice == 4 {
		filename = filepath.FromSlash("./wordlists/4-dice.csv")
	} else {
		filename = filepath.FromSlash("./wordlists/5-dice.csv")
	}

	// Read the corresponding file
	lines, err := csvreader.Read(filename)
	if err != nil {
		panic(err)
	}

	// Loop over file content and find word by ID
	for _, line := range lines {
		if line.ID == n {
			return line.Word
		}
	}

	// No word found for ID `n`
	panic(fmt.Sprintf("Word not found for ID: %v", n))
}

func main() {
	// Use 4 words by default
	words := 4

	// Use n words as indicated in arg
	if len(os.Args) > 1 {
		num, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			panic("Invalid argument(s)")
		}
		words = int(num)
	}

	var pass string
	for i := 0; i < words; i++ {
		if i%2 != 0 {
			// if `i` is odd, shake 4 dice
			n := shaker.Shake(4)
			s := findWord(n, 4)
			pass += s + " "
		} else {
			// if `i` is even, shake 5 dice
			n := shaker.Shake(5)
			s := findWord(n, 5)
			pass += s + " "
		}
	}

	fmt.Println("Your password is:")
	fmt.Printf("\t%v\n", pass)
	fmt.Println()
}
