package main

import (
	"aoc_2022/days"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello AOC")
	argLength := len(os.Args)
	if argLength == 1 {
		fmt.Println("No argument given. Needs name of day given. Aborting.")
		return
	}
	//programName := os.Args[0]
	day := os.Args[1]
	fmt.Printf("day: %v\n", day)
	if day == "01atest" {
		days.Day01TestA()
	} else if day == "01btest" {
		days.Day01TestB()
	} else if day == "01a" {
		days.Day01A()
	} else if day == "01b" {
		days.Day01B()
	}
}
