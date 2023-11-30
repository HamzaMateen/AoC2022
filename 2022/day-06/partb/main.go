// Day 06, Tuning Trouble
// Part 2
// author@HamzaMateen

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	inputFilePath := filepath.Join(currentDir, "./parta/input.txt")

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	charCount := 0
	const NUMBER_OF_CHARS_FOR_MSG = 14
	marker := make([]rune, NUMBER_OF_CHARS_FOR_MSG)

	for { // infinite loop (osm)
		if char, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			// process the rune
			if charCount < NUMBER_OF_CHARS_FOR_MSG {
				marker[charCount] = char
				charCount++
				continue
			}

			// check if all elems are different?
			seen := false
			markerMap := make(map[rune]bool)
			for _, elem := range marker {
				if markerMap[elem] {
					seen = true // Found a duplicate element
				}
				markerMap[elem] = true
			}

			// shifting takes place here
			if seen {
				for i := 1; i < cap(marker); i++ {
					marker[i-1] = marker[i]
				}
				// free to add next rune in the end
				marker[NUMBER_OF_CHARS_FOR_MSG-1] = char
				charCount++

				fmt.Printf("%+v \n", marker)
			} else {
				break
			}
			// read one character successfully
		}
	}

	fmt.Printf("the number of characters till a MSG: %d \n", charCount)
}
