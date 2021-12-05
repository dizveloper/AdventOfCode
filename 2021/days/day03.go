package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	inputs "aoc2021/inputs"
)

// https://adventofcode.com/2021/day/3
// Three : advent of code, day three part1 and 2
func Three() {
	inputSlice := inputs.Day3

	part1(inputSlice)
	part2(inputSlice)
}

func part1(inputSlice []string) {
	gamma := []string{}
	epsilon := []string{}

	for column := 0; column < len(inputSlice[0]); column++ {
		ones := 0
		zeros := 0

		for _, row := range inputSlice {

			runed := []rune(row)

			asNum, _ := strconv.Atoi(string(runed[column]))

			if asNum == 1 {
				ones++
			} else if asNum == 0 {
				zeros++
			}
		}

		if ones > zeros {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}

	gammaBinStr := strings.Join(gamma, "")
	epsilonBinStr := strings.Join(epsilon, "")

	gammaInt, _ := strconv.ParseInt(gammaBinStr, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilonBinStr, 2, 64)

	fmt.Println("Part 1:")
	fmt.Println(gammaInt * epsilonInt)
}

func part2(inputSlice []string) {
	oxygenBin, _ := strconv.ParseInt(p2helper(inputSlice, 1), 2, 64)
	c02Bin, _ := strconv.ParseInt(p2helper(inputSlice, 0), 2, 64)

	fmt.Println("Part 2:")
	fmt.Println(oxygenBin * c02Bin)
}

func p2helper(in []string, digit int) string {
	binStr := ""

	for column := 0; column < len(in[0]); column++ {
		ones := 0
		zeros := 0

		for _, row := range in {

			runed := []rune(row)

			asNum, _ := strconv.Atoi(string(runed[column]))

			if asNum == 1 {
				ones++
			} else if asNum == 0 {
				zeros++
			}
		}

		if len(in) == 1 {
			binStr = in[0]
		} else if len(in) == 2 {
			firstConv, _ := strconv.Atoi(string([]rune(in[0])[column]))
			if firstConv == digit {
				binStr = in[0]
			} else {
				binStr = in[1]
			}
		} else {
			if ones > zeros {
				in = buildArr(digit, column, in)
			} else if zeros > ones {
				temp := float64(digit - 1)
				in = buildArr(int(math.Abs(temp)), column, in)
			} else {
				in = buildArr(digit, column+1, in)
			}
		}
	}

	return binStr
}

func buildArr(digit int, column int, arr []string) []string {
	strippedArr := []string{}

	for _, bin := range arr {
		temp, _ := strconv.Atoi(string([]rune(bin)[column]))

		if temp == digit {
			strippedArr = append(strippedArr, bin)
		}
	}

	return strippedArr
}
