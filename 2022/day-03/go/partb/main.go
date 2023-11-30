package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/juliangruber/go-intersect"
)

func main() {
	// read the input file
	file, err := os.Open("./partb/input.txt")

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

	groupCounter := 0
	groupSlice := make([]string, 3)
	for scanner.Scan() {
		line := scanner.Text()

		// process the line
		groupSlice[groupCounter] = line
		groupCounter++

		if groupCounter == 3 {
			intermediate := intersect.Simple(groupSlice[0], groupSlice[1])
			commonRunes := intersect.Simple(intermediate, groupSlice[2])

			runeValue := rune(commonRunes[0].(uint8))
			prioritySum += priorityMap[runeValue]
			groupCounter = 0
		}
	}

	fmt.Printf("The sum is %d\n", prioritySum)
}
