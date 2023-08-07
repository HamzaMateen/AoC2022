package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// read the file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 0
	totalScore := 0
	for scanner.Scan() {
		count += 1
		line := scanner.Text()

		// process the line
		opponent, me := func() (string, string) {
			x := strings.Split(line, " ")
			return x[0], x[1]
		}()
		// IIFE solution: https://stackoverflow.com/questions/19832189/unpack-slices-on-assignment

		const (
			ROCK     int = 1
			PAPER        = 2
			SCISSORS     = 3
			WIN          = 6
			DRAW         = 3
			LOSS         = 0
		)

		const (
			rock     string = "A"
			paper           = "B"
			scissors        = "C"
		)

		// for PART 1
		switch opponent {
		case rock:
			if me == "X" {
				totalScore += ROCK + DRAW
			} else if me == "Y" {
				totalScore += PAPER + WIN
			} else {
				totalScore += SCISSORS + LOSS
			}
		case paper:
			if me == "X" {
				totalScore += ROCK + LOSS
			} else if me == "Y" {
				totalScore += PAPER + DRAW
			} else {
				totalScore += SCISSORS + WIN
			}
		case scissors:
			if me == "X" {
				totalScore += ROCK + WIN
			} else if me == "Y" {
				totalScore += PAPER + LOSS
			} else {
				totalScore += SCISSORS + DRAW
			}
		}
	}

	// print total according to the strategy guide given
	fmt.Printf("Total %d \n ", totalScore)

}
