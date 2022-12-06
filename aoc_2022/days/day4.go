package days

import (
	"aoc_2022/common"
	"fmt"
	"strings"
)

func Day04A() {
	println("Day 04 A")
	answer := evalDay4A("./res/day4.txt")
	fmt.Println("Num contained is", answer, "Should be 602")
}

func Day04B() {
	println("Day 04 B")
	answer := evalDay4B("./res/day4.txt")
	fmt.Println("Num overlapping is", answer, "Should be 891")
}

func Day04TestA() {
	println("Day 04 test A")
	answer := evalDay4A("./res/day4test.txt")
	fmt.Println("Num contained is", answer, "Should be 2")
}

func Day04TestB() {
	println("Day 04 test B")
	answer := evalDay4B("./res/day4test.txt")
	fmt.Println("Num overlapping is", answer, "Should be 4")

}

func evalDay4A(path string) int {
	lines := common.ReadFileToLines(path)
	count := 0
	for _, line := range lines {
		if isContained(lineToLimits(line)) {
			count++
		}
	}
	return count
}
func evalDay4B(path string) int {
	lines := common.ReadFileToLines(path)
	count := 0
	for _, line := range lines {
		if overlaps(lineToLimits(line)) {
			count++
		}
	}
	return count
}

func isContained(low1, high1, low2, high2 int) bool {
	if low1 <= low2 && high1 >= high2 {
		return true
	}

	if low2 <= low1 && high2 >= high1 {
		return true
	}
	return false
}

func overlaps(low1, high1, low2, high2 int) bool {
	if low1 >= low2 && low1 <= high2 {
		return true
	}
	if low2 >= low1 && low2 <= high1 {
		return true
	}

	if high1 >= low2 && high1 <= high2 {
		return true
	}
	if high2 >= low1 && high2 <= high1 {
		return true
	}

	return false
}

func lineToLimits(line string) (int, int, int, int) {
	sets := strings.Split(line, ",")
	limits1 := strings.Split(sets[0], "-")
	limits2 := strings.Split(sets[1], "-")
	limit1Low := common.StringToInt(limits1[0])
	limit1High := common.StringToInt(limits1[1])
	limit2Low := common.StringToInt(limits2[0])
	limit2High := common.StringToInt(limits2[1])
	return int(limit1Low), int(limit1High), int(limit2Low), int(limit2High)
}
