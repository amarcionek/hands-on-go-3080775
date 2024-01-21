// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	fileName := os.Args[1]

	bs, err := os.ReadFile(fileName)

	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	fileData := string(bs)

	// initialize a map to store the counts
	counts := map[string]int{
		"letters": 0,
		"symbols": 0,
		"numbers": 0,
		"words":   0,
	}

	// use the standard library utility package that can help us split the string into words
	//words := strings.Split(fileData, " ") // This under counted
	words := strings.Fields(fileData)

	// capture the length of the words slice
	numWords := len(words)

	counts["words"] = numWords
	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, char := range word {
			if unicode.IsNumber(char) {
				counts["numbers"]++
			} else if unicode.IsLetter(char) {
				counts["letters"]++
			} else {
				counts["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
