package days

import "fmt"

func RunTask(taskName string) {
	switch taskName {
	case "01atest":
		Day01TestA()
	case "01btest":
		Day01TestB()
	case "01a":
		Day01A()
	case "01b":
		Day01B()
	case "02atest":
		Day02TestA()
	case "02btest":
		Day02TestB()
	case "02a":
		Day02A()
	case "02b":
		Day02B()
	case "03atest":
		Day03TestA()
	case "03btest":
		Day03TestB()
	case "03a":
		Day03A()
	case "03b":
		Day03B()
	case "04atest":
		Day04TestA()
	case "04btest":
		Day04TestB()
	case "04a":
		Day04A()
	case "04b":
		Day04B()
	case "05atest":
		Day05TestA()
	case "05btest":
		Day05TestB()
	case "05a":
		Day05A()
	case "05b":
		Day05B()
	case "06atest":
		Day06TestA()
	case "06btest":
		Day06TestB()
	case "06a":
		Day06A()
	case "06b":
		Day06B()
	case "07atest":
		Day07TestA()
	case "07btest":
		Day07TestB()
	case "07a":
		Day07A()
	case "07b":
		Day07B()
	default:
		fmt.Println("Unknown task given:", taskName)
	}

}
