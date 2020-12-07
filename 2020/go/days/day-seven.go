package days

import (
	"fmt"
	"strconv"
	"strings"

	inputs "../inputs"
)

var bagMap = make(map[string][]string)
var numberOfColorsThatHoldGold = []string{}

// Seven : advent of code, day seven part1 and 2
func Seven() {
	input := inputs.Day7

	inputCleanUp :=
		strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(input, "bags", "", -1), "bag", "", -1), ".", "", -1), " ", "", -1)

	bagLines := strings.Split(inputCleanUp, "\n")

	for line := range bagLines {
		temp := strings.Split(bagLines[line], "contain")
		bagMap[temp[0]] = strings.Split(temp[1], ",")
	}

	colorChecker([]string{"shinygold"})

	fmt.Print("(Part 1) - Total number of colors that can hold a shiny gold bag: ")
	numberOfColorsThatHoldGold = removeDuplicateValues(numberOfColorsThatHoldGold)
	fmt.Println(len(numberOfColorsThatHoldGold))

	fmt.Print("(Part 2) - Total number of bags in a shiny gold one: ")
	bagInBagCalculator("shinygold")

	total := 0
	for col := range totalBagsInBags {
		total = total + totalBagsInBags[col][0]
	}

	fmt.Println(totalBagsInBags)
	fmt.Println(total)
}

func colorChecker(colorsToCheck []string) {
	colorsToCheckAgain := []string{}

	for color := range colorsToCheck {
		for k := range bagMap {
			if doesContain(bagMap[k], colorsToCheck[color]) {
				numberOfColorsThatHoldGold = append(numberOfColorsThatHoldGold, k)
				colorsToCheckAgain = append(colorsToCheckAgain, k)
			}
		}
	}

	colorsToCheckAgain = removeDuplicateValues(colorsToCheckAgain)

	if len(colorsToCheckAgain) > 0 {
		colorChecker(colorsToCheckAgain)
	} else {
		return
	}
}

var totalBagsInBags = make(map[string][]int)

func bagInBagCalculator(colorToStart string) {
	inThatBag := bagMap[colorToStart]

	for within := range inThatBag {
		if !strings.Contains(inThatBag[within], "noother") {
			toRune := []rune(inThatBag[within])
			splitNum, _ := strconv.Atoi(string(toRune[0]))
			color := strings.Replace(inThatBag[within], string(toRune[0]), "", -1)

			totalBagsInBags[color] = []int{splitNum + splitNum*bagsInThatOne(color)}
		}
	}
}

func bagsInThatOne(color string) int {
	totes := 0

	inThatBag := bagMap[color]

	for within := range inThatBag {
		if !strings.Contains(inThatBag[within], "noother") {
			toRune := []rune(inThatBag[within])
			splitNum, _ := strconv.Atoi(string(toRune[0]))
			nextColor := strings.Replace(inThatBag[within], string(toRune[0]), "", -1)

			totes = totes + splitNum + splitNum*bagsInThatOne(nextColor)
		}
	}

	return totes
}

func doesContain(sliceToCheckFrom []string, stringToCheck string) bool {
	for toCheck := range sliceToCheckFrom {
		if strings.Contains(sliceToCheckFrom[toCheck], stringToCheck) {
			return true
		}
	}

	return false
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
