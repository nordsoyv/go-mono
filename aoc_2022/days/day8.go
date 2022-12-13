package days

import "fmt"

func Day08A() {
	println("Day 08 A")
}

func Day08B() {
	println("Day 08 B")
}

func Day08TestA() {
	println("Day 08 test A")
	trees := makeTrees()
	trees.AddRow([]int{0, 1, 2, 3, 4})
	trees.AddRow([]int{5, 6, 7, 8, 9})
	fmt.Println(trees)
	fmt.Println(trees.GetSize())
}

func Day08TestB() {
	println("Day 08 test B")
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
