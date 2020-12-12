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

	snapshot2 := deepCopy2dSlice(seatGrid)
	scannedGrid2 := seatWalker2(seatGrid)
	for !reflect.DeepEqual(snapshot2, scannedGrid2) {
		snapshot2 = deepCopy2dSlice(scannedGrid2)
		tmp := deepCopy2dSlice(scannedGrid2)
		scannedGrid2 = seatWalker2(tmp)
	}

	fmt.Print("(Part2) - Final seat grid with ")
	fmt.Print(walkOccupied(scannedGrid2))
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

func seatWalker2(grid [][]rune) [][]rune {
	newGrid := deepCopy2dSlice(grid)

	for row, cols := range grid {
		for seat := range cols {
			if cols[seat] == 'L' && checkOccupied2(row, seat, grid) == 0 {
				newGrid[row][seat] = '#'
			}

			if cols[seat] == '#' && checkOccupied2(row, seat, grid) >= 5 {
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

	for x := Max(0, row-1); x <= Min(row+1, len(grid)-1); x++ {
		for y := Max(0, seat-1); y <= Min(seat+1, len(grid[0])-1); y++ {
			if !(x == row && y == seat) {
				if grid[x][y] == '#' {
					occupied++
				}
			}
		}
	}
	return occupied
}

func checkOccupied2(row int, seat int, grid [][]rune) int {
	occupied := 0
	/*
		col-1,row-1		col,row-1		col+1, row-1
		col-1,row			X			col+1, row
		col-1,row+1		col,row+1		col+1,row+1
	*/

	for x := Max(0, row-1); x <= Min(row+1, len(grid)-1); x++ {
		for y := Max(0, seat-1); y <= Min(seat+1, len(grid[0])-1); y++ {
			if !(x == row && y == seat) {
				if grid[x][y] == '#' {
					occupied++
				} else if grid[x][y] == '.' {
					occupied = occupied + checkNextNeighbor(grid, x, y, getDir(row, seat, x, y))
				}
			}
		}
	}

	return occupied
}

func getDir(row int, seat int, x int, y int) string {

	if row == x && seat > y {
		return "left"
	}

	if row == x && seat < y {
		return "right"
	}

	if row > x && seat == y {
		return "up"
	}

	if row < x && seat == y {
		return "down"
	}

	if row > x && seat > y {
		return "diagUpLeft"
	}

	if row < x && seat > y {
		return "diagDownLeft"
	}

	if row > x && seat < y {
		return "diagUpRight"
	}

	if row < x && seat < y {
		return "diagDownRight"
	}

	return ""
}

func checkNextNeighbor(grid [][]rune, x int, y int, dir string) int {
	switch dir {
	case "left":
		if y-1 >= 0 {
			if grid[x][y-1] == '.' {
				return checkNextNeighbor(grid, x, y-1, "left")
			} else if grid[x][y-1] == '#' {
				return 1
			}
		}
	case "right":
		if y+1 < len(grid[x]) {
			if grid[x][y+1] == '.' {
				return checkNextNeighbor(grid, x, y+1, "right")
			} else if grid[x][y+1] == '#' {
				return 1
			}
		}
	case "up":
		if x-1 >= 0 {
			if grid[x-1][y] == '.' {
				return checkNextNeighbor(grid, x-1, y, "up")
			} else if grid[x-1][y] == '#' {
				return 1
			}
		}
	case "down":
		if x+1 < len(grid) {
			if grid[x+1][y] == '.' {
				return checkNextNeighbor(grid, x+1, y, "down")
			} else if grid[x+1][y] == '#' {
				return 1
			}
		}
	case "diagUpLeft":
		if x-1 >= 0 && y-1 >= 0 {
			if grid[x-1][y-1] == '.' {
				return checkNextNeighbor(grid, x-1, y-1, "diagUpLeft")
			} else if grid[x-1][y-1] == '#' {
				return 1
			}
		}
	case "diagDownLeft":
		if x+1 < len(grid) && y-1 >= 0 {
			if grid[x+1][y-1] == '.' {
				return checkNextNeighbor(grid, x+1, y-1, "diagDownLeft")
			} else if grid[x+1][y-1] == '#' {
				return 1
			}
		}
	case "diagUpRight":
		if x-1 >= 0 && y+1 < len(grid[x]) {
			if grid[x-1][y+1] == '.' {
				return checkNextNeighbor(grid, x-1, y+1, "diagUpRight")
			} else if grid[x-1][y+1] == '#' {
				return 1
			}
		}
	case "diagDownRight":
		if x+1 < len(grid) && y+1 < len(grid[x]) {
			if grid[x+1][y+1] == '.' {
				return checkNextNeighbor(grid, x+1, y+1, "diagDownRight")
			} else if grid[x+1][y+1] == '#' {
				return 1
			}
		}
	default:
		return 0
	}

	return 0
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
