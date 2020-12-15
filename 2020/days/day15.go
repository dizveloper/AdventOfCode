package days

import (
	"fmt"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/15
// Fifteen : advent of code, day Fifteen part1 and 2
func Fifteen() {
	inputSlice := inputs.Day15

	// k: number spoken | v: {num of times spoken, last time spoken}
	spoken := make(map[int][]int)

	for i, num := range inputSlice {
		spoken[num] = []int{1, i, i}
	}

	fmt.Print("(Part 1) - 2020th number spoken: ")
	fmt.Println(getLastNumberSpoken(spoken, 2020))
	fmt.Print("(Part 1) - 2020th number spoken: ")
	fmt.Println(getLastNumberSpoken(spoken, 30000000))
}

func getLastNumberSpoken(spoken map[int][]int, last int) int {
	nextNum := 0
	var lastNum int
	var lastTurn int
	for _, v := range spoken {
		if v[1] > lastTurn {
			lastTurn = v[1]
			lastNum = v[0]
		}
	}

	for i := len(spoken); i < last; i++ {
		if _, ok := spoken[lastNum]; ok {
			if spoken[lastNum][0] == 1 {
				nextNum = 0
			} else {
				nextNum = spoken[lastNum][1] - spoken[lastNum][2]
			}
		}

		if _, ok := spoken[nextNum]; ok {
			spoken[nextNum][0]++
			spoken[nextNum][2] = spoken[nextNum][1]
			spoken[nextNum][1] = i
			lastNum = nextNum
		} else {
			spoken[nextNum] = []int{1, i, i}
			lastNum = nextNum
		}
	}

	return nextNum
}
