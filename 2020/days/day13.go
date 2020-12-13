package days

import (
	"fmt"
	"strconv"

	inputs "../inputs"
)

// https://adventofcode.com/2020/day/13
// Thirteen : advent of code, day thirteen part1 and 2
func Thirteen() {
	inputSlice := inputs.Day13
	earliestTimeStamp := inputs.Day13EarliestTimestamp
	buses := []int{}

	for bus := range inputSlice {
		if inputSlice[bus] == "x" {
			buses = append(buses, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice[bus])
			buses = append(buses, asInt)
		}
	}

	fmt.Print("(Part 1) - Minutes to catch multiplied by the bus ID: ")
	fmt.Println(catchNextPossibleBus(buses, earliestTimeStamp))

	fmt.Print("(Part 2) - Earliest timestamp of subsequent series: ")
	fmt.Println(startStampOfSubsequentSeries(buses, 0))

}

func catchNextPossibleBus(buses []int, earliestTimeStamp int) int {
	busMap := make(map[int]int)

	for bus := range buses {
		if buses[bus] > 1 {
			busMap[buses[bus]] = buses[bus] - (earliestTimeStamp % buses[bus]) + earliestTimeStamp
		}
	}

	nextBusTimeStamp := 0
	nextBus := 0

	for k, v := range busMap {
		if v < nextBusTimeStamp || nextBusTimeStamp == 0 {
			nextBusTimeStamp = v
			nextBus = k
		}
	}

	return nextBus * (nextBusTimeStamp - earliestTimeStamp)
}

func startStampOfSubsequentSeries(buses []int, multiplier int) int {
	start := 0
	properBusSeries := false

	for !properBusSeries {
		skipFactor := 1
		properBusSeries, skipFactor = checkThisSeries(buses, start, skipFactor)
		if !properBusSeries {
			start = start + skipFactor
		}
	}

	return start
}

func checkThisSeries(buses []int, start int, skipFactor int) (bool, int) {
	for bus := 0; bus < len(buses); bus++ {

		if (start+bus)%buses[bus] != 0 {
			return false, skipFactor
		}

		skipFactor = skipFactor * buses[bus]
	}

	return true, skipFactor
}
