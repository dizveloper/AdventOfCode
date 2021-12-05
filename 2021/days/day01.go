package days

import (
	"fmt"

	inputs "aoc2021/inputs"
)

// https://adventofcode.com/2021/day/1
// One : advent of code, day one part1 and 2
func One() {
	inputSlice := inputs.Day1

	tots := 0

	for i := 1; i < len(inputSlice); i++ {
		if inputSlice[i] > inputSlice[i-1] {
			tots++
		}
	}

	fmt.Println("Part 1:")
	fmt.Println(tots)

	tots2 := 0

	for windowStart := 0; windowStart < len(inputSlice)-3; windowStart++ {
		windowEnd := windowStart + 2

		frontWindowSum := 0
		for i := windowStart; i <= windowEnd; i++ {
			frontWindowSum += inputSlice[i]
		}

		backWindowSum := 0
		for i := windowStart + 1; i <= windowEnd+1; i++ {
			backWindowSum += inputSlice[i]
		}

		if backWindowSum > frontWindowSum {
			tots2++
		}
	}

	fmt.Println("Part 2:")
	fmt.Println(tots2)
}
