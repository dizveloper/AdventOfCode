package days

import (
	"fmt"

	inputs "../inputs"
)

// Three : advent of code, day three part1 and 2.
func Three() {
	inputSlice := inputs.Day3

	slopes := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees := []int{}

	for slope := 0; slope < len(slopes); slope++ {
		treesAtSlope := treesInSlope(inputSlice, slopes[slope][0], slopes[slope][1])
		trees = append(trees, treesAtSlope)
	}

	multiplied := 1
	for tree := 0; tree < len(trees); tree++ {
		multiplied = multiplied * trees[tree]
	}

	fmt.Print("Trees encountered: ")
	fmt.Println(trees)
	fmt.Print("Trees multiplied: ")
	fmt.Println(multiplied)
}

func treesInSlope(grid []string, right int, down int) int {
	trees := 0
	row := 0
	column := 0

	for {
		if column >= len(grid[0]) {
			column = column % len(grid[0])
		}

		if row >= len(grid) {
			break
		}

		if isItATree(grid[row], column) {
			trees++
		}

		column = column + right
		row = row + down
	}

	return trees
}

func isItATree(row string, columnPos int) bool {
	toRune := []rune(row)
	tree := '#'
	return toRune[columnPos] == tree
}
