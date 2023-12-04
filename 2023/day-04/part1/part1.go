// --- Day 4: Scratchcards ---
// Part 1
// solution by HamzaMateen

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func split(delim rune) bool {
	return delim == ':' || delim == '|'
}

func arrayIntersection(arr1, arr2 []int) []int {
	intersection := make([]int, 0)

	// create a set out of the first array
	set := make(map[int]bool)

	for _, val := range arr2 {
		set[val] = true
	}

	for _, val := range arr1 {
		if set[val] {
			intersection = append(intersection, val)
		}
	}

	return intersection
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	inputFilePath := filepath.Join(currentDir, "./input.txt")

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	// var cardID int
	var stringParts []string

	winningNumbers := make([]int, 0)
	elfNumbers := make([]int, 0)

	// var cardId int
	var scratchCardsWorth int = 0

	// read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// split line by : and |
		stringParts = strings.FieldsFunc(line, split)

		// get game id
		cardId, err := strconv.Atoi(strings.Fields(stringParts[0])[1])

		if err != nil {
			log.Fatal("string to int conversion error")
		}

		fmt.Printf("%d\n", cardId)

		// convert winning numbers and user numbers into string of ints
		for _, word := range strings.Fields(stringParts[1]) {
			value, _ := strconv.Atoi(word)
			winningNumbers = append(winningNumbers, value)
		}

		for _, word := range strings.Fields(stringParts[2]) {
			value, _ := strconv.Atoi(word)
			elfNumbers = append(elfNumbers, value)
		}

		// now take the intersection of both arrays of numbers
		intersection := arrayIntersection(winningNumbers, elfNumbers)

		// calculate the cratch cards' worth
		scratchCardsWorth += int(math.Pow(2, float64(len(intersection))-1))

		// empty slices here, or make their capacity 0 again
		winningNumbers = make([]int, 0)
		elfNumbers = make([]int, 0)

	}

	fmt.Printf("The elf's scratchcard's worth is: %d\n", scratchCardsWorth)

}
