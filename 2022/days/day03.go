package days

import (
	"aoc2022/inputs"
	"fmt"
	"strings"
)

// https://adventofcode.com/2022/day/3
// Three : advent of code, day three part1 and 2
func Three() {
	inputSlice := inputs.Day3

	priority_weight := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

	pri_score := 0
	pri_score2 := 0

	for _, rucksack := range inputSlice {
		compartment_size := len(rucksack) / 2

		first_comp := rucksack[:compartment_size]
		second_comp := rucksack[compartment_size:]

		for _, item := range first_comp {
			if strings.Contains(second_comp, string(item)) {
				pri_score += priority_weight[string(item)]
				break
			}
		}
	}

	fmt.Println(pri_score)

	for rucksack := 0; rucksack < len(inputSlice); rucksack += 3 {
		for _, sack := range inputSlice[rucksack] {
			if strings.Contains(inputSlice[rucksack+1], string(sack)) && strings.Contains(inputSlice[rucksack+2], string(sack)) {
				pri_score2 += priority_weight[string(sack)]
				break
			}
		}
	}

	fmt.Println(pri_score2)
}
