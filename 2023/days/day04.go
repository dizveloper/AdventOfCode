package days

import (
	"aoc2022/inputs"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/1
// Four : advent of code, day four part1 and 2

var cardMap = make(map[int]string)
var cardMapV2 = make(map[int][]string)

func Four() {
	cards := inputs.Day4ex

	for _, card := range cards {
		space := regexp.MustCompile(`\s+`)
		card = space.ReplaceAllString(card, " ")

		splitGameAndNumbers := strings.Split(card, ": ")

		card = strings.TrimLeft(splitGameAndNumbers[0], "Card ")
		card = strings.TrimRight(card, ":")
		card = strings.TrimSpace(card)

		hand_number, _ := strconv.Atoi(card)
		cardMap[hand_number] = splitGameAndNumbers[1]
		cardMapV2[hand_number] = []string{splitGameAndNumbers[1]}
	}

	four_part1()
	resMap := four_part2(cardMapV2, 1)
	totalCards := 0
	for i := 0; i <= len(resMap); i++ {
		totalCards += len(resMap[i])
	}

	for k, v := range resMap {
		fmt.Println(k, len(v))
	}

	fmt.Println("Part 2: ", totalCards)
}

func four_part1() {
	scratchPileScore := 0
	for _, card := range cardMap {
		scratchHandScore := 0
		winningCards := strings.Split(strings.Split(card, " | ")[0], " ")
		inHand := strings.Split(card, " | ")[1]

		for _, w := range winningCards {
			w := " " + w + " "
			inHand := " " + inHand + " "
			if strings.Contains(inHand, w) {
				if scratchHandScore == 0 {
					scratchHandScore = 1
				} else {
					scratchHandScore *= 2
				}
			}
		}
		scratchPileScore += scratchHandScore
	}

	fmt.Println("Part 1: ", scratchPileScore)
}

func four_part2(cMap map[int][]string, startingIndex int) map[int][]string {
	for k := startingIndex; k < len(cMap); k++ {
		card := cMap[k]
		scratchHandScore := 0
		winningCards := strings.Split(strings.Split(card[0], " | ")[0], " ")
		inHand := strings.Split(card[0], " | ")[1]

		for i := 0; i < len(card); i++ {
			for _, w := range winningCards {
				w := " " + w + " "
				inHand := " " + inHand + " "
				for i := 0; i < len(card); i++ {
					if strings.Contains(inHand, w) {
						scratchHandScore++
					}
				}
			}
		}

		scratchHandScore *= len(card)

		for i := k + 1; i < scratchHandScore; i++ {
			if i < len(cMap) {
				in := strings.Split(cMap[i][0], " | ")[1]
				cMap[i] = append(cMap[i], in)
			}
		}

		if k == len(cMap) {
			break
		} else {
			cMap = four_part2(cMap, k+1)
		}
	}

	return cMap
}
