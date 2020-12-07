package days

import (
	"fmt"
	"strconv"
	"strings"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/2
// Two : advent of code, day two part1 and 2.
func Two() {
	inputSlice := inputs.Day2

	partOneSlice := []string{}
	partTwoSlice := []string{}

	for i := 0; i < len(inputSlice); i++ {
		rangeSlice := getRange(inputSlice[i])
		ruleAndPassSlice := getRuleAndPass(inputSlice[i])

		occurances := strings.Count(ruleAndPassSlice[1], ruleAndPassSlice[0])
		if occurances >= rangeSlice[0] && occurances <= rangeSlice[1] {
			partOneSlice = append(partOneSlice, inputSlice[i])
		}

		if checkIndices(ruleAndPassSlice[1], []rune(ruleAndPassSlice[0])[0], rangeSlice[0], rangeSlice[1]) {
			partTwoSlice = append(partTwoSlice, inputSlice[i])
		}

	}

	fmt.Print("Part 1: ")
	fmt.Println(len(partOneSlice))

	fmt.Print("Part 2: ")
	fmt.Println(len(partTwoSlice))
}

func checkIndices(in string, r rune, first int, last int) bool {
	toRune := []rune(in)
	return ((toRune[first-1] == r && toRune[last-1] != r) || (toRune[first-1] != r && toRune[last-1] == r))
}

func getRange(input string) []int {
	splitter := strings.Split(input, " ")
	rangeAsStrings := strings.Split(splitter[0], "-")

	rangeAsInts := make([]int, len(rangeAsStrings))

	for i, s := range rangeAsStrings {
		rangeAsInts[i], _ = strconv.Atoi(s)
	}

	return rangeAsInts
}

func getRuleAndPass(input string) []string {
	splitter := strings.Split(input, " ")
	ruleAndPass := []string{strings.Replace(splitter[1], ":", "", -1), splitter[2]}

	return ruleAndPass
}
