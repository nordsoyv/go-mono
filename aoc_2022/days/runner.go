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
	default:
		fmt.Println("Unknown task given:", taskName)
	}

}
