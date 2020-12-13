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

	earliestTimeStampx := 939
	inputSlicex := []string{"7", "13", "x", "x", "59", "x", "31", "19"}
	inputSlice1 := []string{"17", "x", "13", "19"}
	inputSlice2 := []string{"67", "7", "59", "61"}
	inputSlice3 := []string{"67", "x", "7", "59", "61"}
	inputSlice4 := []string{"67", "7", "x", "59", "61"}
	inputSlice5 := []string{"1789", "37", "47", "1889"}
	buses := []int{}
	busesx := []int{}

	for bus := range inputSlice {
		if inputSlice[bus] != "x" {
			asInt, _ := strconv.Atoi(inputSlice[bus])
			buses = append(buses, asInt)
		}
	}

	for bus := range inputSlicex {
		if inputSlicex[bus] != "x" {
			asInt, _ := strconv.Atoi(inputSlicex[bus])
			busesx = append(busesx, asInt)
		}
	}

	fmt.Print("(Part 1) - Minutes to catch multiplied by the bus ID: ")
	fmt.Println(catchNextPossibleBus(buses, earliestTimeStamp))
	fmt.Println(catchNextPossibleBus(busesx, earliestTimeStampx))

	buses = nil
	busesx = nil
	buses1 := []int{}
	buses2 := []int{}
	buses3 := []int{}
	buses4 := []int{}
	buses5 := []int{}

	for bus := range inputSlice {
		if inputSlice[bus] == "x" {
			buses = append(buses, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice[bus])
			buses = append(buses, asInt)
		}
	}

	for bus := range inputSlice1 {
		if inputSlice1[bus] == "x" {
			buses1 = append(buses1, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice1[bus])
			buses1 = append(buses1, asInt)
		}
	}

	for bus := range inputSlice2 {
		if inputSlice2[bus] == "x" {
			buses2 = append(buses2, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice2[bus])
			buses2 = append(buses2, asInt)
		}
	}

	for bus := range inputSlice3 {
		if inputSlice3[bus] == "x" {
			buses3 = append(buses3, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice3[bus])
			buses3 = append(buses3, asInt)
		}
	}

	for bus := range inputSlice4 {
		if inputSlice4[bus] == "x" {
			buses4 = append(buses4, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice4[bus])
			buses4 = append(buses4, asInt)
		}
	}

	for bus := range inputSlice5 {
		if inputSlice5[bus] == "x" {
			buses5 = append(buses5, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlice5[bus])
			buses5 = append(buses5, asInt)
		}
	}

	for bus := range inputSlicex {
		if inputSlicex[bus] == "x" {
			busesx = append(busesx, 1)
		} else {
			asInt, _ := strconv.Atoi(inputSlicex[bus])
			busesx = append(busesx, asInt)
		}
	}

	fmt.Print("(Part 2) - Earliest timestamp of subsequent series: ")
	fmt.Println(startStampOfSubsequentSeries(busesx, 0))
	fmt.Println(startStampOfSubsequentSeries(buses1, 0))
	fmt.Println(startStampOfSubsequentSeries(buses2, 0))
	fmt.Println(startStampOfSubsequentSeries(buses3, 0))
	fmt.Println(startStampOfSubsequentSeries(buses4, 0))
	fmt.Println(startStampOfSubsequentSeries(buses5, 0))
	fmt.Println("ACTUAL:")
	fmt.Println(startStampOfSubsequentSeries(buses, 0))

}

func catchNextPossibleBus(buses []int, earliestTimeStamp int) int {
	busMap := make(map[int]int)

	for bus := range buses {
		busMap[buses[bus]] = buses[bus] - (earliestTimeStamp % buses[bus]) + earliestTimeStamp
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
	busesInSeries := 0
	start := buses[0] * multiplier
	lastCheck := 0
	lenWithoutOnes := 0

	for {
		for bus := 1; bus < len(buses); bus++ {
			if buses[bus] != 1 {
				if (buses[bus]-(start%buses[bus])) <= buses[0] && lastCheck+1 == (buses[bus]-(start%buses[bus])) {
					busesInSeries++
				}
				lenWithoutOnes++
			}
			lastCheck++
		}

		if busesInSeries == lenWithoutOnes {
			break
		}

		multiplier++
		start = buses[0] * (multiplier)
		busesInSeries = 0
		lastCheck = 0
		lenWithoutOnes = 0
	}

	return start
}
