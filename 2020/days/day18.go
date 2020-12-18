package days

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/1
// Eighteen : advent of code, day eighteen part1 and 2
func Eighteen() {
	inputSlice := inputs.Day18

	fmt.Print("(Part 1) - Get sum of all statements: ")
	fmt.Println(getSumOfStatements(inputSlice))

	// NOTE TO SELF: When you get back to this, just put parens around all valid subsets of addition expressions.
	fmt.Println(putParensOnPluses(inputSlice))
	fmt.Print("(Part 2) - Get sum of all statements: ")
	fmt.Println(getSumOfStatements(putParensOnPluses(inputSlice)))
}

func putParensOnPluses(input []string) []string {
	inputModified := []string{}

	return inputModified
}

func getSumOfStatements(statements []string) int {
	bigSum := 0

	for _, statement := range statements {
		bigSum = bigSum + getSumOfStatement(statement)
	}

	return bigSum
}

func getSumOfStatement(statement string) int {
	sum := 0

	asRunes := []rune(statement)

	parens := 0

	for i := 0; i < len(asRunes); i++ {
		if unicode.IsDigit(asRunes[i]) {
			asInt, _ := strconv.Atoi(string(asRunes[i]))
			sum += asInt
		} else if asRunes[i] == '(' {
			parens, sum, i = parenDance(parens, sum, i, asRunes, '+')
		} else if asRunes[i] == '+' {
			i++
			if unicode.IsDigit(asRunes[i]) {
				asInt, _ := strconv.Atoi(string(asRunes[i]))
				sum += asInt
			} else if asRunes[i] == '(' {
				parens, sum, i = parenDance(parens, sum, i, asRunes, '+')
			}
		} else if asRunes[i] == '*' {
			i++
			if unicode.IsDigit(asRunes[i]) {
				asInt, _ := strconv.Atoi(string(asRunes[i]))
				sum *= asInt
			} else if asRunes[i] == '(' {
				parens, sum, i = parenDance(parens, sum, i, asRunes, '*')
			}
		}
	}

	return sum
}

func parenDance(parens int, sum int, i int, asRunes []rune, op rune) (int, int, int) {
	parens++
	str := ""
	for parens > 0 {
		i++
		if asRunes[i] == '(' {
			parens++
		} else if asRunes[i] == ')' {
			parens--
		}

		str = str + string(asRunes[i])

		if parens == 0 {
			str = trimLastChar(str)

			if op == '*' {
				sum *= getSumOfStatement(str)
			} else if op == '+' {
				sum += getSumOfStatement(str)
			}
		}
	}

	return parens, sum, i
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
