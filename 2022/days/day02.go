package days

import (
	"aoc2022/inputs"
	"fmt"
)

// https://adventofcode.com/2021/day/2
// Two : advent of code, day two part1 and 2
func Two() {
	input := inputs.Day2

	score := 0

	score_mappings := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	result_mappings := map[string]int{
		"A X": 3,
		"A Y": 6,
		"A Z": 0,
		//
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		//
		"C X": 6,
		"C Y": 0,
		"C Z": 3,
	}

	for _, round := range input {
		score += (result_mappings[round] + score_mappings[string(round[2])])
	}

	fmt.Println(score)
}
