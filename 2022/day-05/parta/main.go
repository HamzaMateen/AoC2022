// Day 05, Part 1 (CrateMover 9000, picks up only one crate at a time)
// author@HamzaMateen

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
)

func main() {

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

	// create stacks
	const NUM_OF_STACKS = 9
	var stacks [NUM_OF_STACKS][]rune

	// stack order reading loop
	shouldBreakOuterLoop := false
	const CRATE_POSITION_DENOMINATOR = 4
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {

			if line[i] == '1' {
				// read the the empty line and then exit the loop
				scanner.Scan()
				shouldBreakOuterLoop = true
				break
			}

			// process here
			currentStack := int(i / CRATE_POSITION_DENOMINATOR)
			// interesting relation
			if i%4 == 0 {
				if line[i] != ' ' {
					stacks[currentStack] = append(stacks[currentStack], rune(line[i+1]))
				}
			}
		}

		if shouldBreakOuterLoop {
			break
		}
	}

	// since I read the stacks from top-to-bottom. They have to be reversed hence
	for i := 0; i < NUM_OF_STACKS; i++ {
		slices.Reverse(stacks[i])
		var length int = len(stacks[i])
		for j := 0; j < length; j++ {
			fmt.Printf("%s", string(stacks[i][j]))
		}
		fmt.Println()
	}

	// read the move operations from here, let's use regex
	for scanner.Scan() {
		line := scanner.Text()

		pattern := regexp.MustCompile("[0-9]+")

		count, from, to := func() (int, int, int) {
			matches := pattern.FindAllString(line, -1)

			a, _ := strconv.Atoi(matches[0])
			b, _ := strconv.Atoi(matches[1])
			c, _ := strconv.Atoi(matches[2])

			// decrement the indices though
			return a, b - 1, c - 1
		}()

		// plan the moves now, move 1 crate at a time
		startIndex := len(stacks[from]) - count

		removedElements := stacks[from][startIndex:]
		stacks[from] = stacks[from][:startIndex]

		// add to the following slice
		// stacks[to] = append(stacks[to], removedElements...)
		for i := count - 1; i >= 0; i-- {
			stacks[to] = append(stacks[to], removedElements[i])
		}
	}

	// let's print the last runes of each stack
	for i := 0; i < NUM_OF_STACKS; i++ {
		length := len(stacks[i])
		fmt.Printf("%s", string(stacks[i][length-1]))
	}
}
