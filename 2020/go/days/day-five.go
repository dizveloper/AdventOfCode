package days

import (
	"fmt"
	"sort"

	inputs "../inputs"
)

// Five : advent of code, day five part1 and 2
func Five() {
	inputSlice := inputs.Day5

	seatIds := partOne(inputSlice)
	fmt.Println("")
	fmt.Println("Your seat ID is: ")
	fmt.Println(partTwo(seatIds))
}

func partOne(inputSlice []string) []int {
	seatIds := []int{}

	for dir := range inputSlice {
		toRune := []rune(inputSlice[dir])
		rowRange := []int{0, 128}
		colRange := []int{0, 8}
		finalRow := 0
		finalColumn := 0
		for i := 0; i < 7; i++ {
			if toRune[i] == 'F' {
				rowRange[1] = rowRange[1] - ((rowRange[1] - rowRange[0]) / 2)
			} else if toRune[i] == 'B' {
				rowRange[0] = rowRange[0] + ((rowRange[1] - rowRange[0]) / 2)
			}

			if i == 6 {
				finalRow = rowRange[0]
			}
		}

		for i := 0; i < 3; i++ {
			if toRune[i+7] == 'L' {
				colRange[1] = colRange[1] - ((colRange[1] - colRange[0]) / 2)
			} else if toRune[i+7] == 'R' {
				colRange[0] = colRange[0] + ((colRange[1] - colRange[0]) / 2)
			}

			if i == 2 {
				finalColumn = colRange[0]
			}
		}

		seatIds = append(seatIds, ((finalRow * 8) + finalColumn))
	}

	highest := 0
	for i := range seatIds {
		if seatIds[i] > highest {
			highest = seatIds[i]
		}
	}

	fmt.Print("Highest Seat ID: ")
	fmt.Println(highest)
	return seatIds
}

func partTwo(seatIds []int) int {
	sort.Ints(seatIds)
	seatsLeft := []int{}

	for i, seat := 0, seatIds[0]; seat < seatIds[len(seatIds)-1]; i, seat = i+1, seat+1 {
		if seat != seatIds[i] {
			seatsLeft = append(seatsLeft, seat)
			seat = seatIds[i]
		}
	}

	return seatsLeft[0]
}
