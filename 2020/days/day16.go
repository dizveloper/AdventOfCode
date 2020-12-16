package days

import (
	"fmt"
	"strings"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/16
// Sixteen : advent of code, day sixteen part1 and 2
func Sixteen() {

	/*
		This struct is being used for input

		type Tickets struct {
			Fields        map[string][]int
			Ticket        []int
			NearbyTickets [][]int
		}
	*/
	input := inputs.Day16

	fmt.Print("(Part 1) - Sum of all bad field values: ")
	fmt.Println(sumAllInvalidValuesOnNearbyTickets(input))

	withValidTicketsOnly := tossBadTickets(input)
	matchingFieldNames := getFieldNames(withValidTicketsOnly)

	departureTotal := 1

	for i, fieldName := range matchingFieldNames {
		if strings.Contains(fieldName, "departure") {
			departureTotal *= withValidTicketsOnly.Ticket[i]
		}
	}

	fmt.Print("(Part 2) - Multiplied total of all departure related fields on my ticket: ")
	fmt.Println(departureTotal)
}

func sumAllInvalidValuesOnNearbyTickets(tickets inputs.Tickets) int {
	badSum := 0

	for _, ticket := range tickets.NearbyTickets {
		for _, fieldValue := range ticket {
			if !validForAnyField(fieldValue, tickets) {
				badSum = badSum + fieldValue
			}
		}
	}

	return badSum
}

func validForAnyField(fieldValue int, tickets inputs.Tickets) bool {
	isIt := false

	for _, ranges := range tickets.Fields {
		if (fieldValue >= ranges[0] && fieldValue <= ranges[1]) || (fieldValue >= ranges[2] && fieldValue <= ranges[3]) {
			isIt = true
			break
		} else {
			isIt = false
		}
	}

	return isIt
}

func tossBadTickets(tickets inputs.Tickets) inputs.Tickets {
	toReturn := inputs.Tickets{}
	toReturn.Fields = tickets.Fields
	toReturn.Ticket = tickets.Ticket
	toReturn.NearbyTickets = [][]int{}

	for _, ticket := range tickets.NearbyTickets {
		fullTicketMatch := true
		for _, fieldValue := range ticket {
			if !validForAnyField(fieldValue, tickets) {
				fullTicketMatch = false
				break
			}
		}
		if fullTicketMatch {
			toReturn.NearbyTickets = append(toReturn.NearbyTickets, ticket)
		}
	}

	return toReturn
}

func getFieldName(column []int, fields map[string][]int) []string {
	allMatchingFieldNames := []string{}

	for fieldName, ranges := range fields {
		fullMatch := true
		for _, val := range column {
			if !((val >= ranges[0] && val <= ranges[1]) || (val >= ranges[2] && val <= ranges[3])) {
				fullMatch = false
				break
			}
		}

		if fullMatch {
			allMatchingFieldNames = append(allMatchingFieldNames, fieldName)
		}
	}

	return allMatchingFieldNames
}

func getFieldNames(tickets inputs.Tickets) []string {
	columns := [][]int{}
	column := []int{}

	for col := 0; col < len(tickets.NearbyTickets[0]); col++ {
		column = []int{}
		for row := 0; row < len(tickets.NearbyTickets); row++ {
			column = append(column, tickets.NearbyTickets[row][col])
		}
		columns = append(columns, column)
	}

	matchingFieldNames := [][]string{}

	for _, col := range columns {
		matchingFieldNames = append(matchingFieldNames, getFieldName(col, tickets.Fields))
	}

	return filterFieldNames(matchingFieldNames)
}

func filterFieldNames(allMatchingFieldNames [][]string) []string {
	filtered := []string{}

	for !allSizeOne(allMatchingFieldNames) {
		for i, rows := range allMatchingFieldNames {
			for j, col := range rows {

				anotherOfMe := false

				for k, rowsInner := range allMatchingFieldNames {
					for y, colInner := range rowsInner {
						if y != j && k != y {
							if col == colInner {
								anotherOfMe = true
							}
						}
					}
				}

				if !anotherOfMe {
					allMatchingFieldNames[i] = []string{col}
				}
			}
		}
	}

	for _, fieldNameAsSlice := range allMatchingFieldNames {
		filtered = append(filtered, fieldNameAsSlice[0])
	}

	return filtered
}

func allSizeOne(doubleSlice [][]string) bool {
	res := true

	for _, v := range doubleSlice {
		if len(v) > 1 {
			res = false
			break
		}
	}

	return res
}
