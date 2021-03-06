package days

import (
	"fmt"
	"strconv"
	"strings"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/8
// Eight : advent of code, day eight part1 and 2
func Eight() {
	inputSlice := inputs.Day8

	part1(inputSlice)
	part2(inputSlice)
}

func part1(inputSlice []string) {
	acc, _ := oppEvaluator(inputSlice)
	fmt.Print("(Part1) - Accumulator prior to infinite loop restarting: ")
	fmt.Println(acc)
}

var whichOpChanged = []int{0, 0}

func part2(inputSlice []string) {
	innerSlice := inputSlice

	for operation := whichOpChanged[0]; operation < len(innerSlice); operation++ {
		temp := strings.Split(innerSlice[operation], " ")

		typeOfOp := temp[0]
		signedNumber, _ := strconv.Atoi(temp[1])

		if typeOfOp == "nop" && whichOpChanged[0] == operation && whichOpChanged[1] == signedNumber {
			innerSlice[whichOpChanged[0]] = "jmp " + strconv.Itoa(whichOpChanged[1])
		}

		if typeOfOp == "jmp" {
			innerSlice[operation] = "nop " + strconv.Itoa(signedNumber)
			whichOpChanged = []int{operation, signedNumber}
			break
		}
	}

	acc, in := oppEvaluator(innerSlice)
	if in < len(inputSlice) {
		part2(innerSlice)
	} else {
		fmt.Print("(Part2) - Accumulator after termination: ")
		fmt.Println(acc)
	}

}

func oppEvaluator(inputSlice []string) (int, int) {
	accumulator := 0
	operationIndex := 0
	indexVisited := []int{0}

	for {
		temp := strings.Split(inputSlice[operationIndex], " ")

		typeOfOp := temp[0]
		signedNumber, _ := strconv.Atoi(temp[1])

		switch typeOfOp {
		case "acc":
			accumulator = accumulator + signedNumber
			operationIndex++
		case "jmp":
			operationIndex = operationIndex + signedNumber
		case "nop":
			operationIndex++
		default:
			fmt.Println("not an opp")
		}

		if intSliceContains(indexVisited, operationIndex) {
			break
		} else {
			indexVisited = append(indexVisited, operationIndex)
		}

		if operationIndex > len(inputSlice)-1 || operationIndex < 0 {
			break
		}
	}

	return accumulator, operationIndex
}

func intSliceContains(s []int, i int) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
