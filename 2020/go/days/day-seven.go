package days

import (
	"fmt"
	"strings"

	inputs "../inputs"
)

var bagMap = make(map[string]string)
var numberOfColorsThatHoldGold = 0

// Seven : advent of code, day seven part1 and 2
func Seven() {
	input := inputs.Day7
	// 	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
	// dark orange bags contain 3 bright white bags, 4 muted yellow bags.
	// bright white bags contain 1 shiny gold bag.
	// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
	// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
	// dark olive bags contain 3 faded blue bags, 4 dotted black bags.
	// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
	// faded blue bags contain no other bags.
	// dotted black bags contain no other bags.`

	// I could have used NewReplacer here but I already wrote this and it works :)
	inputCleanUp :=
		strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(input, "bags", "", -1), "bag", "", -1), ".", "", -1), " ", "", -1)

	bagLines := strings.Split(inputCleanUp, "\n")

	for line := range bagLines {
		temp := strings.Split(bagLines[line], "contain")
		bagMap[temp[0]] = temp[1]
	}

	colorChecker([]string{"shinygold"}, "")

	fmt.Print("(Part 1) - Total number of colors that can hold a shiny gold bag: ")
	fmt.Println(numberOfColorsThatHoldGold)
}

func colorChecker(colorsToCheck []string, alreadyConsidered string) {
	colorsToCheckAgain := []string{}

	for color := range colorsToCheck {
		for k := range bagMap {
			if strings.Contains(bagMap[k], colorsToCheck[color]) && !strings.Contains(alreadyConsidered, bagMap[k]) {
				numberOfColorsThatHoldGold++
				colorsToCheckAgain = append(colorsToCheckAgain, k)
				alreadyConsidered = alreadyConsidered + bagMap[k]
			}
		}
	}

	colorsToCheckAgain = removeDuplicateValues(colorsToCheckAgain)

	fmt.Println(len(colorsToCheck))
	if len(colorsToCheckAgain) > 0 {
		colorChecker(colorsToCheckAgain, alreadyConsidered)
	} else {
		return
	}
}

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
