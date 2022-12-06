package main

import (
	"fmt"
	"os"

	days "aoc2022/days"
)

func main() {
	whichDay(os.Args[1])
}

func whichDay(day string) {
	switch day {
	case "1":
		fmt.Println("Running day 1")
		days.One()
	case "2":
		fmt.Println("Running day 2")
		days.Two()
	case "3":
		fmt.Println("Running day 3")
		days.Three()
	case "4":
		fmt.Println("Running day 4")
		days.Four()
	case "5":
		fmt.Println("Not there yet.")
	case "6":
		fmt.Println("Not there yet.")
	case "7":
		fmt.Println("Not there yet.")
	case "8":
		fmt.Println("Not there yet.")
	case "9":
		fmt.Println("Not there yet.")
	case "10":
		fmt.Println("Not there yet.")
	case "11":
		fmt.Println("Not there yet.")
	case "12":
		fmt.Println("Not there yet.")
	case "13":
		fmt.Println("Not there yet.")
	case "14":
		fmt.Println("Not there yet.")
	case "15":
		fmt.Println("Not there yet.")
	case "16":
		fmt.Println("Not there yet.")
	case "17":
		fmt.Println("Not there yet.")
	case "18":
		fmt.Println("Not there yet.")
	case "19":
		fmt.Println("Not there yet.")
	case "20":
		fmt.Println("Not there yet.")
	case "21":
		fmt.Println("Not there yet.")
	case "22":
		fmt.Println("Not there yet.")
	case "23":
		fmt.Println("Not there yet.")
	case "24":
		fmt.Println("Not there yet.")
	case "25":
		fmt.Println("Not there yet.")
	default:
		fmt.Println("Pick a day between 1 and 25")
	}
}
