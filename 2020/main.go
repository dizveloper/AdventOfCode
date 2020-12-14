package main

import (
	"fmt"
	"os"

	days "./days"
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
		fmt.Println("Running day 5")
		days.Five()
	case "6":
		fmt.Println("Running day 6")
		days.Six()
	case "7":
		fmt.Println("Running day 7")
		days.Seven()
	case "8":
		fmt.Println("Running day 8")
		days.Eight()
	case "9":
		fmt.Println("Running day 9")
		days.Nine()
	case "10":
		fmt.Println("Running day 10")
		days.Ten()
	case "11":
		fmt.Println("Running day 11")
		days.Eleven()
	case "12":
		fmt.Println("Running day 12")
		days.Twelve()
	case "13":
		fmt.Println("Running day 13")
		days.Thirteen()
	case "14":
		fmt.Println("Running day 14")
		days.Fourteen()
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
