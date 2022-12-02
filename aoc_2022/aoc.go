package main

import (
	"aoc_2022/days"
	"fmt"
	"os"
)

func main() {
	argLength := len(os.Args)
	if argLength == 1 {
		fmt.Println("No argument given. Needs name of day given. Aborting.")
		return
	}
	//programName := os.Args[0]
	day := os.Args[1]
	switch day {
	case "01atest":
		days.Day01TestA()
	case "01btest":
		days.Day01TestB()
	case "01a":
		days.Day01A()
	case "01b":
		days.Day01B()
	case "02atest":
		days.Day02TestA()
	case "02btest":
		days.Day02TestB()
	case "02a":
		days.Day02A()
	case "02b":
		days.Day02B()
	}

}
