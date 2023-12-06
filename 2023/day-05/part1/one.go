// --- Day 5: If You Give A Seed A Fertilizer ---
// Part 1
// solution by HamzaMateen

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getRangeValues(line string) (int, int, int) {
	rangeValues := make([]int, 0)

	for _, word := range strings.Fields(line) {
		value, _ := strconv.Atoi(word)
		rangeValues = append(rangeValues, value)
	}

	return rangeValues[0], rangeValues[1], rangeValues[2]
}

func skipLines(scanner *bufio.Scanner, count int) {
	for i := 0; i < count; i++ {
		scanner.Scan()
		scanner.Text()
	}
}

func populateMap(mapObject *map[int]int, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		dest, src, len := getRangeValues(line)
		for i := 0; i < len; i++ {
			(*mapObject)[src] = dest

			src++
			dest++
		}
	}
}

func getMin(arr []int) int {
	currMin := arr[0]

	for i := 1; i < len(arr); i++ {
		if currMin > arr[i] {
			currMin = arr[i]
		}
	}

	return currMin
}

func main() {
	// reading a file line by line
	filename := "input.txt"
	dirPath, _ := os.Getwd()
	filePath := filepath.Join(dirPath, filename)
	fileHandle, _ := os.Open(filePath)

	scanner := bufio.NewScanner(fileHandle)

	// 1. read the type of seeds and store them into an array
	scanner.Scan() // read first line of input
	line := scanner.Text()

	// read empty line as well
	skipLines(scanner, 1)

	seeds := make([]int, 0)
	for _, word := range strings.Fields(line)[1:] {
		seed, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal("string to int conversion error")
		}
		seeds = append(seeds, seed)
	}

	fmt.Printf("%+v\n", seeds)

	// We need to have 6 maps
	seedToSoil := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&seedToSoil, scanner)

	soilToFertilizer := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&soilToFertilizer, scanner)

	fertilizerToWater := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&fertilizerToWater, scanner)

	waterToLight := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&waterToLight, scanner)

	lightToTemperature := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&lightToTemperature, scanner)

	temperatureToHumidity := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&temperatureToHumidity, scanner)

	humidityToLocation := make(map[int]int)
	skipLines(scanner, 1)
	populateMap(&humidityToLocation, scanner)

	// finding locations now
	locations := make([]int, 0)

	for _, seed := range seeds {
		soil, exists := seedToSoil[seed]
		if !exists {
			soil = seed
		}
		fmt.Printf("%d\t", soil)

		fert, exists := soilToFertilizer[soil]
		if !exists {
			fert = soil
		}
		fmt.Printf("%d\t", fert)

		water, exists := fertilizerToWater[fert]
		if !exists {
			water = fert
		}
		fmt.Printf("%d\t", water)

		light, exists := waterToLight[water]
		if !exists {
			light = water
		}
		fmt.Printf("%d\t", light)

		temp, exists := lightToTemperature[light]
		if !exists {
			temp = light
		}
		fmt.Printf("%d\t", temp)

		humidity, exists := temperatureToHumidity[temp]
		if !exists {
			humidity = temp
		}
		fmt.Printf("%d\t", humidity)

		loc, exists := humidityToLocation[humidity]
		if !exists {
			loc = humidity
		}
		fmt.Printf("%d\t", loc)

		// append location value
		locations = append(locations, loc)

		fmt.Printf("\n")
	}

	fmt.Printf("Lowest location number is: %d\n", getMin(locations))
}
