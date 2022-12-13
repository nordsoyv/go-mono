package days

import (
	"aoc_2022/common"
	"fmt"
)

func Day08A() {
	println("Day 08 A")
	numVisible := taskA("./res/day8.txt")
	fmt.Println("Number of visible trees are", numVisible, "Should be 1812")
}

func Day08B() {
	println("Day 08 B")
	taskB("./res/day8.txt")
}

func Day08TestA() {
	println("Day 08 test A")
	numVisible := taskA("./res/day8test.txt")
	fmt.Println("Number of visible trees are", numVisible, "Should be 21")
}

func Day08TestB() {
	println("Day 08 test B")
	taskBTest("./res/day8test.txt")
}

type Trees struct {
	rows [][]int
}

func makeTrees() Trees {
	rows := make([][]int, 0)
	return Trees{
		rows: rows,
	}
}

func (t *Trees) AddRow(row []int) {
	t.rows = append(t.rows, row)
}

func (t Trees) String() string {
	out := ""
	for _, row := range t.rows {
		for _, height := range row {
			out += fmt.Sprintf("%v", height)
		}
		out += "\n"
	}
	return out
}

func (t Trees) GetSize() (int, int) {
	if len(t.rows) == 0 {
		return 0, 0
	}
	return len(t.rows[0]), len(t.rows)

}

func (t Trees) GetMostScenicTreeScore() int {
	width, height := t.GetSize()
	maxScore := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			score := t.GetScenicScore(x, y)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func (t Trees) GetScenicScore(x, y int) int {
	width, height := t.GetSize()
	scenicLeft := 0
	scenicRight := 0
	scenicTop := 0
	scenicDown := 0

	treeHeight := t.rows[y][x]
	// check north
	for n := y - 1; n >= 0; n-- {
		checkHeight := t.rows[n][x]
		if checkHeight < treeHeight {
			scenicTop++
		}
		if checkHeight >= treeHeight {
			scenicTop++
			break
		}
	}
	// check south
	for n := y + 1; n < height; n++ {
		checkHeight := t.rows[n][x]
		if checkHeight < treeHeight {
			scenicDown++
		}
		if checkHeight >= treeHeight {
			scenicDown++
			break
		}
	}
	// check east
	for n := x + 1; n < width; n++ {
		checkHeight := t.rows[y][n]
		if checkHeight < treeHeight {
			scenicLeft++
		}
		if checkHeight >= treeHeight {
			scenicLeft++
			break
		}
	}
	// check west
	for n := x - 1; n >= 0; n-- {
		checkHeight := t.rows[y][n]
		if checkHeight < treeHeight {
			scenicRight++
		}
		if checkHeight >= treeHeight {
			scenicRight++
			break
		}
	}

	return scenicRight * scenicDown * scenicLeft * scenicTop
}

func (t Trees) IsVisible(x, y int) int {
	width, height := t.GetSize()
	if x == 0 || y == 0 {
		return 1
	}
	if x == width-1 {
		return 1
	}
	if y == height-1 {
		return 1
	}
	treeHeight := t.rows[y][x]
	// check north
	isVisible := true
	for n := y - 1; n >= 0; n-- {
		checkHeight := t.rows[n][x]
		if checkHeight >= treeHeight {
			isVisible = false
		}
	}
	if isVisible {
		return 1
	}
	// check south
	isVisible = true
	for n := y + 1; n < height; n++ {
		checkHeight := t.rows[n][x]
		if checkHeight >= treeHeight {
			isVisible = false
		}
	}
	if isVisible {
		return 1
	}

	// check east
	isVisible = true
	for n := x + 1; n < width; n++ {
		checkHeight := t.rows[y][n]
		if checkHeight >= treeHeight {
			isVisible = false
		}
	}
	if isVisible {
		return 1
	}
	// check west
	isVisible = true
	for n := x - 1; n >= 0; n-- {
		checkHeight := t.rows[y][n]
		if checkHeight >= treeHeight {
			isVisible = false
		}
	}
	if isVisible {
		return 1
	}

	return 0
}

func (t Trees) FindAllVisibleTrees() int {
	sum := 0
	width, height := t.GetSize()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sum += t.IsVisible(x, y)
		}
	}
	return sum
}

func taskA(path string) int {
	lines := common.ReadFileToLines(path)
	numVisible := 0
	trees := makeTrees()
	for _, line := range lines {
		row := make([]int, 0)
		for _, letter := range line {
			row = append(row, int(letter-'0'))
		}
		trees.AddRow(row)
	}
	fmt.Println(trees)
	numVisible = trees.FindAllVisibleTrees()
	return numVisible
}

func taskBTest(path string) {
	lines := common.ReadFileToLines(path)
	trees := makeTrees()
	for _, line := range lines {
		row := make([]int, 0)
		for _, letter := range line {
			row = append(row, int(letter-'0'))
		}
		trees.AddRow(row)
	}
	fmt.Println(trees)
	fmt.Println("Scenic score for tree 2 1 is", trees.GetScenicScore(2, 1), "should be 4")
	fmt.Println("Scenic score for tree 2 1 is", trees.GetScenicScore(2, 3), "should be 8")
}

func taskB(path string) {
	lines := common.ReadFileToLines(path)
	trees := makeTrees()
	for _, line := range lines {
		row := make([]int, 0)
		for _, letter := range line {
			row = append(row, int(letter-'0'))
		}
		trees.AddRow(row)
	}
	fmt.Println(trees)

	fmt.Println("Max scenic score is", trees.GetMostScenicTreeScore(), "should be 315495")
}
