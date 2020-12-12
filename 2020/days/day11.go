package days

import (
	"fmt"
	"reflect"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/11
// Eleven : advent of code, day eleven part1 and 2.
func Eleven() {
	seatGridAsSlices := inputs.Day11
	// seatGridAsSlices := []string{
	// 	"L.LL.LL.LL",
	// 	"LLLLLLL.LL",
	// 	"L.L.L..L..",
	// 	"LLLL.LL.LL",
	// 	"L.LL.LL.LL",
	// 	"L.LLLLL.LL",
	// 	"..L.L.....",
	// 	"LLLLLLLLLL",
	// 	"L.LLLLLL.L",
	// 	"L.LLLLL.LL",
	// }

	seatGrid := [][]rune{}
	for seatRow := range seatGridAsSlices {
		asRunes := []rune(seatGridAsSlices[seatRow])
		seatGrid = append(seatGrid, asRunes)
	}

	snapshot := deepCopy2dSlice(seatGrid)
	scannedGrid := seatWalker(seatGrid)
	for !reflect.DeepEqual(snapshot, scannedGrid) {
		snapshot = deepCopy2dSlice(scannedGrid)
		tmp := deepCopy2dSlice(scannedGrid)
		scannedGrid = seatWalker(tmp)
	}

	fmt.Print("(Part1) - Final seat grid with ")
	fmt.Print(walkOccupied(scannedGrid))
	fmt.Println(" occupied.")

}

func seatWalker(grid [][]rune) [][]rune {
	newGrid := deepCopy2dSlice(grid)

	for row, cols := range grid {
		for seat := range cols {
			if cols[seat] == 'L' && checkOccupied(row, seat, grid) == 0 {
				newGrid[row][seat] = '#'
			}

			if cols[seat] == '#' && checkOccupied(row, seat, grid) >= 4 {
				newGrid[row][seat] = 'L'
			}
		}
	}

	return newGrid
}

func walkOccupied(grid [][]rune) int {
	occupied := 0

	for _, cols := range grid {
		for seat := range cols {
			if cols[seat] == '#' {
				occupied++
			}
		}
	}

	return occupied
}

func checkOccupied(row int, seat int, grid [][]rune) int {
	occupied := 0
	/*
		col-1,row-1		col,row-1		col+1, row-1
		col-1,row			X			col+1, row
		col-1,row+1		col,row+1		col+1,row+1
	*/

	if len(grid) >= 0 {
		for x := Max(0, row-1); x <= Min(row+1, len(grid)-1); x++ {
			for y := Max(0, seat-1); y <= Min(seat+1, len(grid[0])-1); y++ {
				if x != row || y != seat {
					if grid[x][y] == '#' {
						occupied++
					}
				}
			}
		}
	}
	return occupied
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func prettyGrid(seatGrid [][]rune) {
	for runes := range seatGrid {
		for rune := range seatGrid[runes] {
			fmt.Print(string(seatGrid[runes][rune]))
		}
		fmt.Println()
	}
}

func deepCopy2dSlice(twoDimensionalSlice [][]rune) [][]rune {
	copySlice := make([][]rune, len(twoDimensionalSlice))
	for i := range twoDimensionalSlice {
		copySlice[i] = make([]rune, len(twoDimensionalSlice[i]))
		copy(copySlice[i], twoDimensionalSlice[i])
	}

	return copySlice
}
