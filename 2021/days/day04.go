package days

import (
	"fmt"
	"strconv"

	inputs "aoc2021/inputs"
)

// https://adventofcode.com/2021/day/4
// Four : advent of code, day four part1 and 2
func Four() {
	calls := inputs.Day4Calls
	boards := inputs.Day04Input

out:
	for _, call := range calls {
		for board := 0; board < len(boards); board++ {
			boards[board] = markBoard(boards[board], call)

			bingo, winningBoard := bingoValidator(boards[board])

			if bingo {
				unmarkedSum := getUnmarkedBoardSum(winningBoard)

				callAsNum, _ := strconv.Atoi(call)
				fmt.Println("Part 1:")
				fmt.Println(unmarkedSum * callAsNum)
				break out
			}
		}
	}

	// fmt.Println("Part 2:")
	// fmt.Println()
}

func markBoard(board [][]string, call string) [][]string {
	tempBoard := board

	for row := 0; row < len(tempBoard); row++ {
		for el := 0; el < len(tempBoard[row]); el++ {
			if tempBoard[row][el] == call {
				tempBoard[row][el] = "x"
			}
		}
	}

	return tempBoard
}

func bingoValidator(board [][]string) (bool, [][]string) {
	bingRow := checkRows(board)
	bingCol := checkCols(board)

	if bingCol || bingRow {
		return true, board
	} else {
		return false, [][]string{}
	}
}

func checkRows(board [][]string) bool {
	xCheck := 0

out:
	for _, row := range board {
		for _, el := range row {
			if el == "x" {
				xCheck++
			}

			if xCheck == 5 {
				return true
				break out
			}
		}
		xCheck = 0
	}

	return false
}

func checkCols(board [][]string) bool {
	xCheck := 0

out:
	for col := 0; col < len(board[0]); col++ {
		for _, row := range board {
			if row[col] == "x" {
				xCheck++
			}

			if xCheck == 5 {
				return true
				break out
			}
		}
		xCheck = 0
	}

	return false
}

func getUnmarkedBoardSum(board [][]string) int {
	sum := 0

	for _, row := range board {
		for _, el := range row {
			if el != "x" {
				asInt, _ := strconv.Atoi(el)
				sum += asInt
			}
		}
	}

	return sum
}
