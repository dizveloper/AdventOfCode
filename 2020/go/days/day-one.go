package days

import (
	"fmt"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/1
// One : advent of code, day one part1 and 2
func One() {
	inputSlice := inputs.Day1

	for i := 0; i < len(inputSlice); i++ {
		for j := i + 1; j < len(inputSlice); j++ {
			for y := j + 1; y < len(inputSlice); y++ {
				if (inputSlice[i] + inputSlice[j] + inputSlice[y]) == 2020 {
					fmt.Println(inputSlice[i] * inputSlice[j] * inputSlice[y])
				}
			}
		}
	}
}
