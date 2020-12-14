package days

import (
	"fmt"
	"strconv"
	"strings"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/14
// Fourteen : advent of code, day fourteen part1 and 2
func Fourteen() {
	inputSlice := inputs.Day14

	// inputSlice := []string{
	// 	"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
	// 	"mem[8] = 11",
	// 	"mem[7] = 101",
	// 	"mem[8] = 0",
	// }

	twoD := [][]string{}

	for _, line := range inputSlice {
		line := strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "mem[", "", -1)
		line = strings.Replace(line, "]", "", -1)
		split := strings.Split(line, "=")
		twoD = append(twoD, split)
	}

	fmt.Print("(Part1) - Bitmask leftovers sum: ")
	fmt.Println(masqueradeLeftiesSum(twoD))
	fmt.Println("(Part2) - I gave up (:")
}

func masqueradeLeftiesSum(instructions [][]string) int {
	var mask string
	mem := map[int]int{}
	sum := 0
	for _, inst := range instructions {
		if inst[0] == "mask" {
			mask = inst[1]
		} else {
			memAddr, _ := strconv.Atoi(inst[0])
			memVal, _ := strconv.Atoi(inst[1])

			flipOnes := strings.ReplaceAll(mask, "X", "1")
			flipZeros := strings.ReplaceAll(mask, "X", "0")

			and, _ := strconv.ParseInt(flipOnes, 2, 0)
			or, _ := strconv.ParseInt(flipZeros, 2, 0)

			andOR := memVal&int(and) | int(or)

			sum = sum + andOR - mem[memAddr]
			mem[memAddr] = andOR
		}
	}
	return sum
}
