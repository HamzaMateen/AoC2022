package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	inputFilePath := filepath.Join(currentDir, "./partb/input.txt")

	// load the file
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}

	// define a scanner
	scanner := bufio.NewScanner(file)
	totalOverlappingSections := 0

	// scan the file
	for scanner.Scan() {
		line := scanner.Text()
		// process the line here
		// chop by ,
		elfOne, elfTwo := func() (string, string) {
			sections := strings.Split(line, ",")
			return sections[0], sections[1]
		}()

		// extract ranges
		rangeOneStart, rangeOneEnd := func() (int, int) {
			_range := strings.Split(elfOne, "-")

			start, _ := strconv.Atoi(_range[0])
			end, _ := strconv.Atoi(_range[1])

			return start, end
		}()

		rangeTwoStart, rangeTwoEnd := func() (int, int) {
			_range := strings.Split(elfTwo, "-")

			start, _ := strconv.Atoi(_range[0])
			end, _ := strconv.Atoi(_range[1])

			return start, end
		}()

		// Part 2 : See if there are partial overlaps

		if rangeOneStart >= rangeTwoStart && rangeOneStart <= rangeTwoEnd {
			totalOverlappingSections += 1
		} else if rangeTwoStart >= rangeOneStart && rangeTwoStart <= rangeOneEnd {
			totalOverlappingSections += 1
		}

	}

	// print the number of overlapping ranges
	fmt.Printf("The overlapping sections' count is: %d \n", totalOverlappingSections)
}
