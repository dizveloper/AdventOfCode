package days

import (
	"aoc2022/inputs"
	"fmt"
	"strings"
)

// https://adventofcode.com/2022/day/5
// Six : advent of code, day six part1 and 2
func Six() {

	marker_window_size := 4
	msg_marker_window_size := 14

	run(marker_window_size)
	run(msg_marker_window_size)
}

func run(window_size int) {
	signal := inputs.Day6

	for let := window_size; let < len(signal); let++ {

		repeats := false
		parsed_sig := signal[let-window_size : let]
		for i := window_size - 1; i > 0; i-- {
			if strings.Count(parsed_sig, string(parsed_sig[i])) > 1 {
				repeats = true
				break
			}
		}

		if !repeats {
			fmt.Println(let)
			break
		}
	}
}
