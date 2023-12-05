package days

import (
	"aoc2022/inputs"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

// https://adventofcode.com/2022/day/1
// Three : advent of code, day three part1 and 2

var engineMap = inputs.Day3

func Three() {

	three_part1()
	three_part2()
}

func three_part1() {
	/*
		Read each line of the engineMap	until you reach a number
		then check all adjacent characters for a symbol
	*/

	partsTotal := 0

	for x := 0; x < len(engineMap); x++ {
		recordCurrent := ""
		for y := 0; y < len(engineMap[x]); y++ {
			if unicode.IsNumber(rune(engineMap[x][y])) {
				recordCurrent += string(engineMap[x][y])
				if y == len(engineMap[x])-1 {
					if _, err := strconv.Atoi(recordCurrent); err == nil && containsAdjacentSymbol(x, y-len(recordCurrent), y) {
						asInt, _ := strconv.Atoi(recordCurrent)
						partsTotal += asInt
					}
					recordCurrent = ""
				}
			} else {
				if _, err := strconv.Atoi(recordCurrent); err == nil && containsAdjacentSymbol(x, y-len(recordCurrent), y) {
					asInt, _ := strconv.Atoi(recordCurrent)
					partsTotal += asInt
				}
				recordCurrent = ""
			}
		}

	}
	fmt.Println("Parts total: ", partsTotal)
}

func containsAdjacentSymbol(x int, startY int, endY int) bool {
	r, _ := regexp.Compile("[^0-9.]")
	/*
		Skip
		x-1 if x = 0
		x+1 if x = len(engineMap)
		y-1 if startY = 0
		y+1 if endY = len(engineMap[x])
	*/
	skipAbove := x == 0
	skipBelow := x == len(engineMap)-1
	skipLeft := startY == 0
	skipRight := endY == len(engineMap[x])-1

	leftOffset := 1
	if skipLeft {
		leftOffset = 0
	}

	rightOffset := 1
	if skipRight {
		rightOffset = 0
	}

	res := false

	if r.MatchString(string(engineMap[x][startY-leftOffset : endY+rightOffset])) {
		res = true
	}

	if !skipAbove {
		if r.MatchString(string(engineMap[x-1][startY-leftOffset : endY+rightOffset])) {
			res = true
		}
	}

	if !skipBelow {
		if r.MatchString(string(engineMap[x+1][startY-leftOffset : endY+rightOffset])) {
			res = true
		}
	}

	return res
}

func three_part2() {
	gearsTotal := 0

	var reggy = regexp.MustCompile(`\d+`)
	var checkIndices = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	var gearsChecked = map[Gear][]string{}

	for k, x := range engineMap {
		var indices = reggy.FindAllIndex([]byte(x), -1)
		for _, i := range indices {
			var start, end = i[0], i[1]
			for y := start; y < end; y++ {
				for _, c := range checkIndices {
					if isSymbol(engineMap, y+c[0], k+c[1]) && engineMap[k+c[1]][y+c[0]] == '*' {
						gearsChecked[Gear{y + c[0], k + c[1]}] = append(gearsChecked[Gear{y + c[0], k + c[1]}], string(engineMap[k][start:end]))
					}
				}
			}
		}
	}

	for _, gear := range gearsChecked {
		var num1 = gear[0]
		var num2 = ""
		for i := range gear {
			if num1 != gear[i] {
				num2 = gear[i]
				break
			}
		}
		if num2 != "" {
			var n1, _ = strconv.Atoi(num1)
			var n2, _ = strconv.Atoi(num2)
			gearsTotal += n1 * n2
		}
	}

	fmt.Println("Gears total: ", gearsTotal)
}

func isSymbol(input []string, x, y int) bool {
	if x >= 0 && y >= 0 && len(input) > y && len(input[y]) > x {
		return input[y][x] != '.' && !unicode.IsDigit(rune(input[y][x]))
	} else {
		return false
	}
}

type Gear struct {
	x, y int
}

// 78236071
