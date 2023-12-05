package days

import (
	"aoc2022/inputs"
	"fmt"
	"slices"
)

// https://adventofcode.com/2022/day/1
// Five : advent of code, day five part1 and 2
var seedMappings = inputs.Day5

func Five() {
	five_part1()
}

func five_part1() {
	seeds := seedMappings["seeds"][0]
	lowestLocation := 999999999999999999
	for _, seed := range seeds {

		// Refactor to kick off a goroutine for each of these and wait till they all complete to get the lowset location

		soil := getter(seed, "seed-to-soil")
		fertilizer := getter(soil, "soil-to-fertilizer")
		water := getter(fertilizer, "fertilizer-to-water")
		light := getter(water, "water-to-light")
		temp := getter(light, "light-to-temperature")
		humidity := getter(temp, "temperature-to-humidity")
		location := getter(humidity, "humidity-to-location")

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println("Part 1: ", lowestLocation)
}

func getter(something int, someKey string) int {
	fmt.Println(someKey)
	someMap := seedMappings[someKey]
	dest := something
	for _, some := range someMap {
		dest = findDestination(some, something)
		if dest != something {
			break
		}
	}
	return dest
}

func findDestination(mappings []int, source int) int {
	sourceWindow := makeWindow(mappings[1], mappings[2], true, source)
	destWindow := makeWindow(mappings[0], mappings[2], false, 0)

	sourceIndex := slices.Index(sourceWindow, source)

	if sourceIndex >= 0 && sourceIndex <= len(destWindow) {
		return destWindow[sourceIndex]
	}

	return source
}

func makeWindow(start int, size int, shortCircuit bool, source int) []int {
	window := []int{}
	for i := 0; i < size; i++ {
		window = append(window, start+i)

		if shortCircuit && start+i == source {
			break
		}
	}
	return window
}

func five_part2() {

}
