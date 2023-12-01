package days

import (
	"aoc2022/inputs"
	"fmt"
	"strconv"
	"unicode"
)

// https://adventofcode.com/2022/day/1
// One : advent of code, day one part1 and 2
func One() {
	calib_vals := inputs.Day1

	calib := ""
	sumCalib := 0

	for _, val := range calib_vals {
		asRunes := []rune(val)
		for r := 0; r < len(asRunes); r++ {
			if unicode.IsDigit(asRunes[r]) {
				calib += string(asRunes[r])
				for r := len(asRunes) - 1; r >= 0; r-- {
					if unicode.IsDigit(asRunes[r]) {
						calib += string(asRunes[r])
						asInt, _ := strconv.Atoi(calib)
						sumCalib += asInt
						break
					}
				}
				calib = ""
				break
			}
		}
	}
	fmt.Println("Part1:", sumCalib)

	// listOfNumbersAsWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// for _, val := range calib_vals {

	// }

	fmt.Println("Part2:")
}
