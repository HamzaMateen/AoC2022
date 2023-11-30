package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// read the input file
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	inputFilePath := filepath.Join(currentDir, "./parta/input.txt")

	file, err := os.Open(inputFilePath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// create a priority map for the alphabets given
	priorityMap := make(map[rune]int)

	for r := 'a'; r <= 'z'; r++ { // 'a''s unicode value is 97, so priority becomes (97-97 + 1)
		priorityMap[r] = int(r - 'a' + 1)
	}

	for r := 'A'; r <= 'Z'; r++ { // 'A' unicode value is 65.
		priorityMap[r] = int(r - 'A' + 27)
	}

	prioritySum := 0
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)

		// process the line
		portionOne := line[0 : length/2]
		portionTwo := line[(length / 2):length]

		// Problem: Find the item type that appears in both compartments of each rucksack.
		for _, char := range portionOne {
			index := strings.IndexRune(portionTwo, char)

			if index > -1 {
				prioritySum += priorityMap[char]

				// break after first finding
				break
			}
		}
	}

	fmt.Printf("The sum is %d\n", prioritySum)
}
