package days

import (
	"aoc2022/inputs"
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/1
// Two : advent of code, day two part1 and 2
var gameMap = make(map[int]string)

func Two() {
	dice_games := inputs.Day2

	// Parse dice_games till ':' remove strip "Game " convert remaining to int
	for _, game := range dice_games {
		splitGameAndNumbers := strings.Split(game, ":")

		game = strings.TrimLeft(splitGameAndNumbers[0], "Game ")
		game = strings.TrimRight(game, ":")

		game_number, _ := strconv.Atoi(game)
		gameMap[game_number] = splitGameAndNumbers[1]
	}

	part1()
	part2()
}

func part1() {
	allowed_dice := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	idTotal := 0
	rollsAllowed := true

	for id, game := range gameMap {
		// split games by ';' to get each roll
		rolls := strings.Split(game, ";")

		// split  each roll by ',' to get each die
		rollsAllowed = true
		for _, roll := range rolls {
			dice := strings.Split(roll, ",")
			for _, die_roll := range dice {
				die_roll := strings.Trim(die_roll, " ")
				die := strings.Split(die_roll, " ")
				// check if die is allowed
				dieNum, _ := strconv.Atoi(die[0])
				if allowed_dice[die[1]] < dieNum {
					rollsAllowed = false
					break
				}
			}
			if !rollsAllowed {
				break
			}
		}

		if rollsAllowed {
			idTotal += id
		}
	}
	fmt.Println("Total of possible game IDs:", idTotal)
}

func part2() {
	allGamesPowerSums := 0
	for _, game := range gameMap {
		setPowerSum := 1
		powerMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		// split games by ';' to get each roll
		rolls := strings.Split(game, ";")

		// split  each roll by ',' to get each die
		for _, roll := range rolls {
			dice := strings.Split(roll, ",")
			for _, die_roll := range dice {
				die_roll := strings.Trim(die_roll, " ")
				die := strings.Split(die_roll, " ")

				dieNum, _ := strconv.Atoi(die[0])

				if powerMap[die[1]] < dieNum {
					powerMap[die[1]] = dieNum
				}
			}
		}

		for _, power := range powerMap {
			if power != 0 {
				setPowerSum = power * setPowerSum
			}
		}

		allGamesPowerSums += setPowerSum
		setPowerSum = 1
		powerMap = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
	}
	fmt.Println("Sum of set powers:", allGamesPowerSums)
}
