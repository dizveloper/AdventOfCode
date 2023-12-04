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

func three_part2() {
	/*
		Read each line of the engineMap	until you reach a number
		then check all adjacent characters for a symbol
	*/

	gearsTotal := 0

	for x := 0; x < len(engineMap); x++ {
		for y := 0; y < len(engineMap[x]); y++ {
			if string(engineMap[x][y]) == "*" {
				// if containsTwoAdjacentNums(x, y) {
				// fmt.Println("Found two adjacent numbers at ", x, y)
				// asInt, _ := strconv.Atoi(string(engineMap[x][y-1 : y+1]))
				// gearsTotal *= asInt
				// }
				fmt.Println("Gear total at:", containsTwoAdjacentNums(x, y))
				gearsTotal += containsTwoAdjacentNums(x, y)
				break
			}
		}

	}
	fmt.Println("Gears total: ", gearsTotal)
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

func containsTwoAdjacentNums(x int, y int) int {
	r, _ := regexp.Compile(`\d`)

	skipAbove := x == 0
	skipBelow := x == len(engineMap)-1
	skipLeft := y == 0
	skipRight := y == len(engineMap[x])-1

	leftOffset := 1
	if skipLeft {
		leftOffset = 0
	}

	rightOffset := 2
	if skipRight {
		rightOffset = 1
	}

	numAbove := 1
	numBelow := 1
	numNextTo := 1
	numNextToLeft := 1
	numNextToRight := 1
	pass := 0

	if !skipAbove {
		// fmt.Println(string(engineMap[x-1][y-leftOffset : y+rightOffset]))
		if r.MatchString(string(engineMap[x-1][y-leftOffset : y+rightOffset])) {
			pass += 1
			numAbove, _ = strconv.Atoi(buildStringNum(x-1, y, leftOffset, rightOffset))
			fmt.Println("Num above: ", numAbove)
		}
	}

	// fmt.Println(string(engineMap[x][y-leftOffset : y+rightOffset]))
	if r.MatchString(string(engineMap[x][y-leftOffset : y+rightOffset])) {
		numNextTo, _ = strconv.Atoi(buildStringNum(x, y, leftOffset, rightOffset))
		fmt.Println("Num next to: ", numNextTo)

		if rune(engineMap[x][y]) == '*' {
			if unicode.IsDigit(rune(engineMap[x][y-leftOffset])) {
				pass += 1
				numNextToLeft, _ = strconv.Atoi(lookLeft(x, y-leftOffset))
				fmt.Println("Num next to Left: ", numNextToLeft)
				// fmt.Println(lookLeft(x, y-leftOffset))
			}

			if unicode.IsDigit(rune(engineMap[x][y+rightOffset-1])) {
				pass += 1
				numNextToRight, _ = strconv.Atoi(lookRight(x, y+rightOffset-1))
				fmt.Println("Num next to Right: ", numNextToRight)
				// fmt.Println(lookRight(x, y+rightOffset))
			}
		}

		if buildStringNum(x, y, leftOffset, rightOffset) == strconv.Itoa(numNextToLeft)+strconv.Itoa(numNextToRight) {
			numNextTo = 1
		}
	}

	if !skipBelow {
		// fmt.Println(string(engineMap[x+1][y-leftOffset : y+rightOffset]))
		if r.MatchString(string(engineMap[x+1][y-leftOffset : y+rightOffset])) {
			pass += 1
			numBelow, _ = strconv.Atoi(buildStringNum(x+1, y, leftOffset, rightOffset))
			fmt.Println("Num below: ", numBelow)
		}
	}

	if pass < 2 {
		return 0
	}
	return numAbove * numBelow * numNextTo * numNextToLeft * numNextToRight
}

func buildStringNum(x int, y int, leftOffset int, rightOffset int) string {
	num := ""
	rightOffset -= 1

	// fmt.Println(engineMap[x][y-leftOffset:y+rightOffset])

	if unicode.IsDigit(rune(engineMap[x][y-leftOffset])) {
		num += lookLeft(x, y-leftOffset)
		// fmt.Println(lookLeft(x, y-leftOffset))
	}
	if unicode.IsDigit(rune(engineMap[x][y])) {
		num += string(engineMap[x][y])
		// fmt.Println(string(engineMap[x][y]))
	}
	if unicode.IsDigit(rune(engineMap[x][y+rightOffset])) {
		num += lookRight(x, y+rightOffset)
		// fmt.Println(lookRight(x, y+rightOffset))
	}

	// fmt.Println("Num: ", num)

	return num
}

func lookLeft(x int, y int) string {
	builder := ""
	for l := y; l >= 0; l-- {
		if unicode.IsDigit(rune(engineMap[x][l])) {
			builder = string(engineMap[x][l]) + builder
		} else {
			break
		}
	}
	return builder
}

func lookRight(x int, y int) string {
	builder := ""
	for r := y; r < len(engineMap[x]); r++ {
		if unicode.IsDigit(rune(engineMap[x][r])) {
			builder += string(engineMap[x][r])
		} else {
			break
		}
	}
	return builder
}
