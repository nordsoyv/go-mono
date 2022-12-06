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
	days.RunTask(day)
}
