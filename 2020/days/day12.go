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

	for dir := range directions {
		asInt, _ := strconv.Atoi(directions[dir][1])
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
			switch asInt {
			case 90:
				// find current in compass and move down 1
				for way := range compass {
					if compass[way] == currentDirection {
						if way-1 < 0 {
							tmp := len(compass) - absOfInt(way-1)
							currentDirection = compass[tmp]
							break
						} else {
							currentDirection = compass[way-1]
							break
						}
					}
				}
			case 180:
				// find current in compass and move down 2
				for way := range compass {
					if compass[way] == currentDirection {
						if way-2 < 0 {
							tmp := len(compass) - absOfInt(way-2)
							currentDirection = compass[tmp]
							break
						} else {
							currentDirection = compass[way-2]
							break
						}
					}
				}
			case 270:
				// find current in compass and move down 3
				for way := range compass {
					if compass[way] == currentDirection {
						if way-3 < 0 {
							tmp := len(compass) - absOfInt(way-3)
							currentDirection = compass[tmp]
							break
						} else {
							currentDirection = compass[way-3]
							break
						}
					}
				}
			default:
				fmt.Println("shrug1.")
			}
		case "R":
			switch asInt {
			case 90:
				// find current in compass and move down 1
				for way := range compass {
					if compass[way] == currentDirection {
						if way+1 >= len(compass) {
							tmp := way + 1 - len(compass)
							currentDirection = compass[tmp]
							break
						} else {
							currentDirection = compass[way+1]
							break
						}
					}
				}
			case 180:
				// find current in compass and move down 2
				for way := range compass {
					if compass[way] == currentDirection {
						if way+2 >= len(compass) {
							tmp := way + 2 - len(compass)
							currentDirection = compass[tmp]
							break
						} else {
							currentDirection = compass[way+2]
							break
						}
					}
				}
			case 270:
				// find current in compass and move down 3
				for way := range compass {
					if compass[way] == currentDirection {
						if way+3 >= len(compass) {
							tmp := way + 3 - len(compass)
							currentDirection = compass[tmp]
							break
						} else {
							currentDirection = compass[way+3]
							break
						}
					}
				}
			default:
				fmt.Println("shrug2.")
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
				fmt.Println("shrug3.")
			}
		default:
			fmt.Println("shrug4.")
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
