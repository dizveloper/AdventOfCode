package days

import (
	"aoc2022/inputs"
	"fmt"
	"strconv"
	"strings"
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

	mapOfNums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	sumCalib2 := 0

	for _, val := range calib_vals {
		extracted := map[int]int{} // K = index found, V = value

		for k, v := range mapOfNums {
			if strings.Contains(val, k) {
				strInd := strings.Index(val, k)
				strIndLast := strings.LastIndex(val, k)
				extracted[strInd] = v
				extracted[strIndLast] = v
			}
		}

		comboStr := strconv.Itoa(extracted[minKey(extracted)]) + strconv.Itoa(extracted[maxKey(extracted)])
		comboInt, _ := strconv.Atoi(comboStr)
		sumCalib2 += comboInt
	}

	fmt.Println("Part2: ", sumCalib2)
}

func maxKey(numbers map[int]int) int {
	max := -10000
	for k := range numbers {
		if k > max {
			max = k
		}
	}
	return max
}

func minKey(numbers map[int]int) int {
	min := 100000000000000000
	for k := range numbers {
		if k < min {
			min = k
		}
	}
	return min
}
