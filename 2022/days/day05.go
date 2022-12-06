package days

import (
	"aoc2022/inputs"
	"fmt"
)

// https://adventofcode.com/2022/day/5
// Five : advent of code, day five part1 and 2
func Five() {
	loads := inputs.Day5_a
	loads2 := make(map[int]string)
	for k, v := range loads {
		loads2[k] = v
	}

	instructions := inputs.Day5_b

	for _, inst := range instructions {
		amount := inst[0]
		from := inst[1]
		to := inst[2]

		for i := 0; i < amount; i++ {
			loads[to] = loads[to] + string(loads[from][len(loads[from])-1])
			loads[from] = loads[from][:len(loads[from])-1]
		}
	}

	fmt.Print(string(loads[1][len(loads[1])-1]))
	fmt.Print(string(loads[2][len(loads[2])-1]))
	fmt.Print(string(loads[3][len(loads[3])-1]))
	fmt.Print(string(loads[4][len(loads[4])-1]))
	fmt.Print(string(loads[5][len(loads[5])-1]))
	fmt.Print(string(loads[6][len(loads[6])-1]))
	fmt.Print(string(loads[7][len(loads[7])-1]))
	fmt.Print(string(loads[8][len(loads[8])-1]))
	fmt.Print(string(loads[9][len(loads[9])-1]))
	fmt.Println()

	for _, inst := range instructions {
		amount := inst[0]
		from := inst[1]
		to := inst[2]

		loads2[to] = loads2[to] + string(loads2[from][len(loads2[from])-amount:])
		loads2[from] = loads2[from][:len(loads2[from])-amount]
	}

	fmt.Print(string(loads2[1][len(loads2[1])-1]))
	fmt.Print(string(loads2[2][len(loads2[2])-1]))
	fmt.Print(string(loads2[3][len(loads2[3])-1]))
	fmt.Print(string(loads2[4][len(loads2[4])-1]))
	fmt.Print(string(loads2[5][len(loads2[5])-1]))
	fmt.Print(string(loads2[6][len(loads2[6])-1]))
	fmt.Print(string(loads2[7][len(loads2[7])-1]))
	fmt.Print(string(loads2[8][len(loads2[8])-1]))
	fmt.Print(string(loads2[9][len(loads2[9])-1]))
	fmt.Println()
}
