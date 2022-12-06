package days

import (
	"fmt"
	"sort"

	"aoc2022/inputs"
)

// https://adventofcode.com/2022/day/1
// One : advent of code, day one part1 and 2
func One() {
	all_calories := inputs.Day1

	totals_by_elf := []int{}

	for _, elf_cals := range all_calories {
		totals_by_elf = append(totals_by_elf, sum(elf_cals))
	}

	sort.Ints(totals_by_elf)

	fmt.Print("highest: ")
	fmt.Println(totals_by_elf[len(totals_by_elf)-1])

	fmt.Print("3 highest total: ")
	fmt.Println(totals_by_elf[len(totals_by_elf)-1] + totals_by_elf[len(totals_by_elf)-2] + totals_by_elf[len(totals_by_elf)-3])
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
