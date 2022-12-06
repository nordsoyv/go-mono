package days

import (
	"aoc_2022/common"
	"fmt"
	"strings"
	"unicode"
)

func Day03A() {
	println("Day 03 A")
	sum := evalA("./res/day3.txt")
	fmt.Println("Sum priority is ", sum, " Should be 8139")

}

func Day03B() {
	println("Day 03 B")
	sum := evalB("./res/day3.txt")
	fmt.Println("Sum priority is ", sum, " Should be 2668")
}

func Day03TestA() {
	println("Day 03 test A")
	sum := evalA("./res/day3test.txt")
	fmt.Println("Sum priority is ", sum, " Should be 157")
}

func Day03TestB() {
	println("Day 03 test B")
	sum := evalB("./res/day3test.txt")
	fmt.Println("Sum priority is ", sum, " Should be 70")
}

func evalA(path string) int {
	lines := common.ReadFileToLines(path)
	sumPriority := 0
	for _, line := range lines {
		lineLength := len(line)
		left := line[:lineLength/2]
		right := line[lineLength/2:]
		for _, leftChar := range []rune(left) {
			if strings.ContainsRune(right, leftChar) {
				sumPriority += getPriority(leftChar)
				break
			}
		}
	}
	return sumPriority
}

func evalB(path string) int {
	lines := common.ReadFileToLines(path)
	elfTrios := chunkSlice(lines, 3)
	fmt.Println(len(elfTrios))
	sumPriority := 0
	for _, trio := range elfTrios {
		for _, char := range []rune(trio[0]) {
			if strings.ContainsRune(trio[1], char) && strings.ContainsRune(trio[2], char) {
				sumPriority += getPriority(char)
				break
			}
		}
	}
	return sumPriority
}

func getPriority(char rune) int {
	if unicode.IsUpper(char) {
		return int(char - 'A' + 26 + 1)
	} else {
		return int(char - 'a' + 1)
	}
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
