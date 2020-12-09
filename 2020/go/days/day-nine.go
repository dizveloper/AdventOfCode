package days

import (
	"fmt"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/9
// Nine : advent of code, day nine part1 and 2
func Nine() {
	inputSlice := inputs.Day9
	// inputSlice := []int{
	// 	35,
	// 	20,
	// 	15,
	// 	25,
	// 	47,
	// 	40,
	// 	62,
	// 	55,
	// 	65,
	// 	95,
	// 	102,
	// 	117,
	// 	150,
	// 	182,
	// 	127,
	// 	219,
	// 	299,
	// 	277,
	// 	309,
	// 	576,
	// }

	preambleSize := 25 //25

	for i := preambleSize; i < len(inputSlice); i++ {
		preamble := inputSlice[i-preambleSize : i]
		if !preambleEquator(preamble, inputSlice[i]) {
			fmt.Print("(Part 1) - Value that doesn't follow preamble rule: ")
			fmt.Println(inputSlice[i])
		}
	}
}

func preambleEquator(preamble []int, current int) bool {
	worked := false
	for j := 0; j < len(preamble); j++ {
		for x := j + 1; x < len(preamble); x++ {
			if preamble[j]+preamble[x] == current {
				worked = true
				break
			}
		}
	}
	return worked
}
