package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	// read the file
// 	file, err := os.Open("./input.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	scanner := bufio.NewScanner(file)

// 	count := 0
// 	totalScore := 0
// 	for scanner.Scan() {
// 		count += 1
// 		line := scanner.Text()

// 		// process the line
// 		opponent, me := func() (string, string) {
// 			x := strings.Split(line, " ")
// 			return x[0], x[1]
// 		}()
// 		// IIFE solution: https://stackoverflow.com/questions/19832189/unpack-slices-on-assignment

// 		const (
// 			X    int = 1
// 			Y        = 2
// 			Z        = 3
// 			WIN      = 6
// 			DRAW     = 3
// 			LOSS     = 0
// 		)

// 		// for PART 2
// 		switch opponent {
// 		case "A": // ROCK
// 			if me == "X" { // gotta lose
// 				totalScore += Z + LOSS
// 			} else if me == "Y" {
// 				totalScore += X + DRAW
// 			} else {
// 				totalScore += Y + WIN
// 			}
// 		case "B": // PAPER
// 			if me == "X" {
// 				totalScore += X + LOSS
// 			} else if me == "Y" {
// 				totalScore += Y + DRAW
// 			} else {
// 				totalScore += Z + WIN
// 			}
// 		case "C": // SCISSORS
// 			if me == "X" {
// 				totalScore += Y + LOSS
// 			} else if me == "Y" {
// 				totalScore += Z + DRAW
// 			} else {
// 				totalScore += X + WIN
// 			}
// 		}
// 	}

// 	// print total according to the strategy guide given
// 	fmt.Printf("Total %d \n ", totalScore)

// }
