package days

import (
	"aoc2022/inputs"
	"fmt"
)

// https://adventofcode.com/2022/day/1
// Five : advent of code, day five part1 and 2
var seedMappings = inputs.Day5

func Five() {
	fmt.Println("Part 1: ", five_part1())
	five_part2()
	fmt.Println("Part 2: ", five_part1())
}

func five_part1() int {
	seeds := seedMappings["seeds"][0]
	lowestLocation := 999999999999999999
	for _, seed := range seeds {
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

	return lowestLocation
}

func getter(something int, someKey string) int {
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
	correspondingDiff := 0
	inRange := false
	if source >= mappings[1] && source <= mappings[1]+mappings[2]-1 {
		correspondingDiff = source - mappings[1]
		inRange = true
	}

	destinationCorresponding := mappings[0] + correspondingDiff

	if correspondingDiff == 0 && !inRange {
		return source
	} else {
		return destinationCorresponding
	}
}

func five_part2() {
	// Expand seeds (this is a massive time sync but smooth afterwards)
	builder := []int{}

	for i := 0; i < len(seedMappings["seeds"][0]); i = i + 2 {
		for j := seedMappings["seeds"][0][i]; j < seedMappings["seeds"][0][i]+seedMappings["seeds"][0][i+1]; j++ {
			fmt.Println(j)
			builder = append(builder, j)
		}
	}
	seedMappings["seeds"][0] = builder
}
