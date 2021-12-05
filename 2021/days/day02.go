package days

import (
	"fmt"
	"strconv"
	"strings"

	inputs "aoc2021/inputs"
)

// https://adventofcode.com/2021/day/2
// Two : advent of code, day two part1 and 2
func Two() {
	inputSlice := inputs.Day2

	horizontal := 0
	vertical := 0

	for _, instruction := range inputSlice {
		direction := strings.Split(instruction, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

		switch direction {
		case "forward":
			horizontal += steps
		case "up":
			vertical -= steps
		case "down":
			vertical += steps
		default:
			fmt.Println("Wrong direction.")
		}
	}

	fmt.Println("Part 1:")
	fmt.Println(horizontal * vertical)

	horizontal = 0
	vertical = 0
	aim := 0

	for _, instruction := range inputSlice {
		direction := strings.Split(instruction, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

		switch direction {
		case "forward":
			horizontal += steps
			vertical += aim * steps
		case "up":
			aim -= steps
		case "down":
			aim += steps
		default:
			fmt.Println("Wrong direction.")
		}
	}

	fmt.Println("Part 2:")
	fmt.Println(horizontal * vertical)
}
