package days

import (
	"fmt"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/9
// Nine : advent of code, day nine part1 and 2
func Nine() {
	inputSlice := inputs.Day9

	part1 := getRuleBreaker(inputSlice)
	fmt.Print("(Part 1) - Value that doesn't follow preamble rule: ")
	fmt.Println(part1)

	part2 := contiguousCheckSum(inputSlice, part1)
	fmt.Print("(Part 2) - Contiguous series sum (series max+min): ")
	fmt.Println(part2)
}

func getRuleBreaker(inputSlice []int) int {
	preambleSize := 25

	for i := preambleSize; i < len(inputSlice); i++ {
		preamble := inputSlice[i-preambleSize : i]
		if !preambleEquator(preamble, inputSlice[i]) {
			return inputSlice[i]
		}
	}

	return 0
}

func contiguousCheckSum(inputSlice []int, toSum int) int {
	maxMinSum := 0

	for i := 0; i < len(inputSlice); i++ {
		for j := i + 1; j < len(inputSlice); j++ {
			if toSum != inputSlice[i] && toSum != inputSlice[j] {
				contiguousSeries := inputSlice[i:j]

				contSum := 0
				for c := 0; c < len(contiguousSeries); c++ {
					contSum = contSum + contiguousSeries[c]
				}

				if contSum == toSum {
					maxMinSum = maxMinSumation(contiguousSeries)
				}
			}
		}
	}

	return maxMinSum
}

func maxMinSumation(contiguousSeries []int) int {
	min := contiguousSeries[0]
	for v := 0; v < len(contiguousSeries); v++ {
		if contiguousSeries[v] < min {
			min = contiguousSeries[v]
		}
	}

	max := contiguousSeries[0]
	for v := 0; v < len(contiguousSeries); v++ {
		if contiguousSeries[v] > max {
			max = contiguousSeries[v]
		}
	}

	return min + max
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
