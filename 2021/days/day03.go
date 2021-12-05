package days

import (
	"fmt"
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
	oxygenBinStr := ""
	c02BinStr := ""

	copy1 := inputSlice
	copy2 := inputSlice

	// oxygen := []string{}
	// c02 := []string{}

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

		if len(copy1) == 1 {
			oxygenBinStr = copy1[0]
		} else if len(copy1) == 2 {
			firstConv, _ := strconv.Atoi(string([]rune(copy1[0])[column+1]))
			if firstConv == 1 {
				oxygenBinStr = copy1[0]
			} else {
				oxygenBinStr = copy1[1]
			}
		} else {
			if ones > zeros {
				copy1 = buildArr(1, column, copy1)
				fmt.Print("copy1 in first: ")
				fmt.Println(copy1)
			} else if zeros > ones {
				copy1 = buildArr(0, column, copy1)
				fmt.Print("copy1 in second: ")
				fmt.Println(copy1)
			} else {
				copy1 = buildArr(1, column, copy1)
				fmt.Print("copy1 in ...: ")
				fmt.Println(copy1)
			}
		}

		if len(copy2) == 1 {
			c02BinStr = copy2[0]
		} else if len(copy2) == 2 {
			firstConv, _ := strconv.Atoi(string([]rune(copy2[0])[column+1]))
			if firstConv == 0 {
				c02BinStr = copy2[0]
			} else {
				c02BinStr = copy2[1]
			}
		} else {
			if ones > zeros {
				copy2 = buildArr(0, column, copy2)
				fmt.Print("copy2 in first: ")
				fmt.Println(copy2)
			} else if zeros > ones {
				copy2 = buildArr(1, column, copy2)
				fmt.Print("copy2 in second: ")
				fmt.Println(copy2)
			} else {
				copy2 = buildArr(0, column, copy2)
				fmt.Print("copy2 in ...: ")
				fmt.Println(copy2)
			}
		}
	}

	oxygenBin, _ := strconv.ParseInt("011011000110", 2, 64)
	c02Bin, _ := strconv.ParseInt("100100101100", 2, 64)

	fmt.Println("Part 2:")
	fmt.Println(oxygenBin * c02Bin)
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
