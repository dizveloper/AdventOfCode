package days

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/4
// Four : advent of code, day four part1 and 2
func Four() {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"ecl",
		"pid",
		"hcl",
		"cid",
	}

	input := inputs.Day4

	passports := strings.Split(input, "\n\n")
	valid := 0
	completeMatch := false

	for i := 0; i < len(passports); i++ {
		for field := 0; field < len(requiredFields); field++ {
			if requiredFields[field] == "cid" {
				continue
			}

			if !strings.Contains(passports[i], requiredFields[field]) {
				completeMatch = false
				break
			} else if !isFieldValid(requiredFields[field], passports[i]) {
				completeMatch = false
				break
			}
			completeMatch = true
		}
		if completeMatch {
			valid++
			completeMatch = false
		}
	}
	fmt.Print("Number of valid passports reported: ")
	fmt.Println(valid)
}

func isFieldValid(field string, passport string) bool {
	passport = strings.Replace(passport, "\n", " ", -1)
	fields := strings.Split(passport, " ")
	for fieldIn := 0; fieldIn < len(fields); fieldIn++ {
		fields[fieldIn] = strings.Replace(fields[fieldIn], "\t", "", -1)
		if strings.Contains(fields[fieldIn], field) {
			fields[fieldIn] = strings.Split(fields[fieldIn], ":")[1]
			switch field {
			case "byr":
				return byr(fields[fieldIn])
			case "iyr":
				return iyr(fields[fieldIn])
			case "eyr":
				return eyr(fields[fieldIn])
			case "hgt":
				return hgt(fields[fieldIn])
			case "ecl":
				return ecl(fields[fieldIn])
			case "pid":
				return pid(fields[fieldIn])
			case "hcl":
				return hcl(fields[fieldIn])
			case "cid":
				return true
			default:
				fmt.Println("not a field")
			}
		}
	}
	return false
}

func hgt(in string) bool {
	// a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.

	if strings.Contains(in, "cm") {
		in = strings.ReplaceAll(in, "cm", "")
		asInt, _ := strconv.Atoi(in)
		return asInt >= 150 && asInt <= 193
	} else if strings.Contains(in, "in") {
		in = strings.ReplaceAll(in, "in", "")
		asInt, _ := strconv.Atoi(in)
		return asInt >= 59 && asInt <= 76
	} else {
		return false
	}
}

func byr(in string) bool {
	// four digits; at least 1920 and at most 2002.
	asInt, _ := strconv.Atoi(in)
	return asInt >= 1920 && asInt <= 2002
}

func iyr(in string) bool {
	// four digits; at least 2010 and at most 2020.
	asInt, _ := strconv.Atoi(in)
	return asInt >= 2010 && asInt <= 2020
}

func eyr(in string) bool {
	// four digits; at least 2020 and at most 2030.
	asInt, _ := strconv.Atoi(in)
	return asInt >= 2020 && asInt <= 2030
}
func hcl(in string) bool {
	// a # followed by exactly six characters 0-9 or a-f.
	allGood := false
	if strings.HasPrefix(in, "#") {
		in = strings.ReplaceAll(in, "#", "")
		if !(len(in) == 6) {
			return false
		}
		toRune := []rune(in)

		for r := 0; r < len(toRune); r++ {
			allGood = (unicode.IsDigit(toRune[r]) || unicode.IsLetter(toRune[r]))
		}
	}
	return allGood
}

func ecl(in string) bool {
	// exactly one of: amb blu brn gry grn hzl oth.
	correctVals := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	for _, val := range correctVals {
		if val == in {
			return true
		}
	}
	return false
}

func pid(in string) bool {
	// a nine-digit number, including leading zeroes.
	if !(len(in) == 9) {
		return false
	}
	_, err := strconv.Atoi(in)
	return err == nil
}
