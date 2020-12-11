package days

import (
	"fmt"
	"sort"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/10
// Ten : advent of code, day ten part1 and 2
func Ten() {
	inputSlice := inputs.Day10

	// Adding in built-ins and sorting
	inputSlice = append(inputSlice, 0)
	sort.Ints(inputSlice)
	inputSlice = append(inputSlice, inputSlice[len(inputSlice)-1]+3)

	fmt.Print("(Part1) - result from multipling the 1 and 3 jolt jumps in the total adapter path: ")
	fmt.Println(numberOfOneAndThreeJoltJumpsInAdapterPath(inputSlice))

	fmt.Print("(Part2) - Number of adapter paths that result in the max joltage of is : ")
	fmt.Println(evaluateAdapterPath(inputSlice, 0))
}

func numberOfOneAndThreeJoltJumpsInAdapterPath(adapterSlice []int) int {
	ones := 0
	threes := 0

	for joltage := 0; joltage < len(adapterSlice)-1; joltage++ {
		if adapterSlice[joltage+1]-adapterSlice[joltage] == 1 {
			ones++
		} else if adapterSlice[joltage+1]-adapterSlice[joltage] == 3 {
			threes++
		}
	}
	return ones * threes
}

var adapterPathsTried = make(map[int]int)

func evaluateAdapterPath(adaptersTocheck []int, startAdapter int) int {
	workingAdapterPaths := 0
	nextAdapter := startAdapter + 1
	furthestPossibleAdapter := startAdapter + 3

	if furthestPossibleAdapter >= len(adaptersTocheck) {
		return 1
	} else if _, ok := adapterPathsTried[adaptersTocheck[startAdapter]]; ok {
		return adapterPathsTried[adaptersTocheck[startAdapter]]
	}

	for i := nextAdapter; i <= furthestPossibleAdapter; i++ {
		lower := adaptersTocheck[startAdapter]
		upper := adaptersTocheck[i]

		if (upper-lower) >= 1 && (upper-lower) <= 3 {
			workingAdapterPaths += evaluateAdapterPath(adaptersTocheck, i)
		}
	}

	adapterPathsTried[adaptersTocheck[startAdapter]] = workingAdapterPaths
	return workingAdapterPaths
}
