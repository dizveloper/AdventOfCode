package days

import (
	"fmt"
	"strconv"
	"strings"
	// inputs "../inputs"
)

// https://adventofcode.com/2020/day/12
// Twelve : advent of code, day twelve part1 and 2
func Twelve() {
	// inputSlice := inputs.Day12
	inputSlice := []string{
		"F 10",
		"N 3",
		"F 7",
		"R 90",
		"F 11",
	}

	directions := [][]string{}
	directions = append(directions, []string{"E", "0"})
	for dir := range inputSlice {
		tmp := strings.Split(inputSlice[dir], " ")
		directions = append(directions, tmp)
	}

	// startingDirection := "E"
	// startingPosition := "0"

	eastWest, northSouth := travel(directions)
	fmt.Print("(Part1) - Manhattan Distance of my ship: ")
	fmt.Println(absOfInt(eastWest) + absOfInt(northSouth))
}

func travel(directions [][]string) (int, int) {
	eastWest := 0
	northSouth := 0
	compass := []string{
		"N",
		"E",
		"S",
		"W",
	}
	currentDirection := directions[0][0]
	// currentPosition := directions[0][1]
	asInt, _ := strconv.Atoi(directions[0][1])

	for dir := range directions {
		switch directions[dir][0] {
		case "N":
			northSouth = northSouth + asInt
		case "S":
			northSouth = northSouth - asInt
		case "E":
			eastWest = eastWest + asInt
		case "W":
			eastWest = eastWest - asInt
		case "L":
			// column % len(grid[0]) (mod to account for over or underflow)
			switch asInt {
			case 90:
				// find current in compass and move down 1
				for way := range compass {
					if compass[way] == currentDirection {
						currentDirection = compass[way-1]
					}
				}
				currentDirection = ""
			case 180:
				// find current in compass and move down 2
				for way := range compass {
					if compass[way] == currentDirection {
						currentDirection = compass[way-2]
					}
				}
				currentDirection = ""
			case 270:
				// find current in compass and move down 3
				for way := range compass {
					if compass[way] == currentDirection {
						currentDirection = compass[way-3]
					}
				}
				currentDirection = ""
			default:
				fmt.Println("shrug.")
			}
		case "R":
			switch asInt {
			case 90:
				// find current in compass and move up 1
				currentDirection = ""
			case 180:
				// find current in compass and move up 2
				currentDirection = ""
			case 270:
				// find current in compass and move up 3
				currentDirection = ""
			default:
				fmt.Println("shrug.")
			}
		case "F":
			switch currentDirection {
			case "N":
				northSouth = northSouth + asInt
			case "S":
				northSouth = northSouth - asInt
			case "E":
				eastWest = eastWest + asInt
			case "W":
				eastWest = eastWest - asInt
			default:
				fmt.Println("shrug.")
			}
		default:
			fmt.Println("shrug.")
		}
	}

	return eastWest, northSouth
}

func absOfInt(val int) int {
	if val < 0 {
		val = -val
	}

	return val
}
