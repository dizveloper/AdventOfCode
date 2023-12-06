package days

import (
	"aoc2022/inputs"
	"fmt"
)

// https://adventofcode.com/2022/day/1
// Six : advent of code, day six part1 and 2

type race struct {
	timeInMs              int
	distanceTraveledToWin int
	amountOfWaysToWin     int
}

func Six() {
	fmt.Println("Day 6:")

	fmt.Println("Part 1: ", calculate(inputs.Day6p1))
	fmt.Println("Part 2: ", calculate(inputs.Day6p2))
}

func calculate(input map[string][]int) int {
	races := []race{}

	for i := 0; i < len(input["Time"]); i++ {
		currentRace := &race{
			timeInMs:              input["Time"][i],
			distanceTraveledToWin: input["Distance"][i],
			amountOfWaysToWin:     0,
		}
		setAmountOfWaysToWin(currentRace)
		races = append(races, *currentRace)
	}

	multiplied := 1
	for _, race := range races {
		multiplied *= race.amountOfWaysToWin
	}

	return multiplied
}

func setAmountOfWaysToWin(r *race) {
	for timeHeld := 1; timeHeld < r.timeInMs; timeHeld++ {
		if timeHeld*(r.timeInMs-timeHeld) > r.distanceTraveledToWin {
			r.amountOfWaysToWin++
		}
	}
}
