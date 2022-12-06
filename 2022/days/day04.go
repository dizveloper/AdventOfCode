package days

import (
	"aoc2022/inputs"
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/4
// Four : advent of code, day four part1 and 2
func Four() {
	inputSlice := inputs.Day4

	full_overlaps := 0
	full_overlaps2 := 0

	for _, elf_pair := range inputSlice {
		elf_one_split := strings.Split(elf_pair[0], "-")

		elf_one_start, _ := strconv.Atoi(elf_one_split[0])
		elf_one_end, _ := strconv.Atoi(elf_one_split[1])

		elf_two_split := strings.Split(elf_pair[1], "-")

		elf_two_start, _ := strconv.Atoi(elf_two_split[0])
		elf_two_end, _ := strconv.Atoi(elf_two_split[1])

		if (elf_one_start <= elf_two_start && elf_one_end >= elf_two_end) || (elf_two_start <= elf_one_start && elf_two_end >= elf_one_end) {
			full_overlaps++
		}

		if (elf_one_start <= elf_two_start && elf_one_end >= elf_two_end) ||
			(elf_one_start <= elf_two_start && elf_one_end <= elf_two_end && elf_one_end >= elf_two_start) ||
			(elf_two_start <= elf_one_start && elf_two_end >= elf_one_end) ||
			(elf_two_start <= elf_one_start && elf_two_end <= elf_one_end && elf_two_end >= elf_one_start) {
			full_overlaps2++
		}
	}

	fmt.Println(full_overlaps)
	fmt.Println(full_overlaps2)
}
