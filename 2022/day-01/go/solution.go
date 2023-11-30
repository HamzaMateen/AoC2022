package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	// so i divide the problem as follows

	// 0. Create an array that would contain the calorie *sum* carried by each elf
	// -> this part gives me how I can deal with dynamic number of entries in an array! (dynamic arr)
	var calorieSums []float64

	// 1. read the puzzle input line by line from a file! (save input to a file)
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Print(err)
	}
	defer file.Close() // discouraged for writing file streams, but we are just reading the file so
	// defer will work

	sum := 0.0
	max := 0.0
	// 2. parseToInt and sum until you hit the '\n' empty line.
	scanner := bufio.NewScanner(file)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) != "" {
			floatNum, err := strconv.ParseFloat(line, 64) // 64 bit float

			if err != nil {
				panic(err)
			}

			sum += floatNum

		} else {
			// 3. push the sum to the end of array
			calorieSums = append(calorieSums, sum)

			// 4. meanwhile keep the 'max' variable to track the maximum sum yet
			if sum > max {
				max = sum
			}

			sum = 0
		}
	}

	// 5. that's the answer to the part 1
	fmt.Printf("Answer 1: %.0f \n", max)

	// Part 2
	// 1. Just sort in ascending order
	sort.Slice(calorieSums, func(i, j int) bool {
		// sort in descending order
		return calorieSums[i] > calorieSums[j]
	})

	// max 3
	sum = calorieSums[0] + calorieSums[1] + calorieSums[2]
	// 2. first 3 indices summed up are the answer!! Kaboom <3
	fmt.Printf("Answer 2: %.0f \n", sum)
}
