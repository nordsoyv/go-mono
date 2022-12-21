package days

import (
	"aoc_2022/common"
	"fmt"
)

func Day12A() {
	println("Day 12 A")
	task12a("./res/day12.txt", -1)
}

func Day12B() {
	println("Day 12 B")
	task12b("", -1)
}

func Day12TestA() {
	println("Day 12 test A")
	task12a("./res/day12test.txt", -1)
}

func Day12TestB() {
	println("Day 12 test B")
	task12b("", -1)
}

func task12a(path string, expected int) {
	lines := common.ReadFileToLines(path)
	m := makeMap(lines)
	fmt.Println(m)
}
func task12b(path string, expected int) {

}

type Map struct {
	start coord
	end   coord
	data  [][]int32
}

func makeMap(lines []string) Map {
	m := Map{
		start: coord{
			x: -1,
			y: -1,
		},
		end: coord{
			x: -1,
			y: -1,
		},
		data: make([][]int32, 0),
	}
	m.readMap(lines)
	return m
}

func (m *Map) readMap(input []string) {
	for y, line := range input {
		row := make([]int32, 0)
		for x, tile := range line {
			if tile == 83 {
				row = append(row, 0)
				m.start.x = x
				m.start.y = y
				continue
			}
			if tile == 69 {
				row = append(row, 123-97)
				m.end.x = x
				m.end.y = y
				continue
			}
			row = append(row, tile-97)
		}
		m.data = append(m.data, row)
	}
}

func (m Map) String() string {
	res := fmt.Sprintf("Start is (%3d,%3d)\n", m.start.x, m.start.y)
	res += fmt.Sprintf("End is   (%3d,%3d)\n", m.end.x, m.end.y)

	for _, row := range m.data {
		for _, tile := range row {
			res += fmt.Sprintf("%2d ", tile)
		}
		res += "\n"
	}
	return res
}

func (m Map) getTileHeight(x, y int) int32 {
	if y < 0 || y >= len(m.data) {
		return 1000
	}
	if x < 0 || x >= len(m.data[y]) {
		return 1000
	}
	return m.data[y][x]
}

type coord struct {
	x, y int
}

func (m Map) getLegalMoves(x, y int) []coord {
	coords := make([]coord, 0)
	currentHeight := m.getTileHeight(x, y)
	if m.getTileHeight(x+1, y) == currentHeight+1 || m.getTileHeight(x+1, y) <= currentHeight {
		coords = append(coords, coord{
			x: x + 1,
			y: y,
		})
	}
	if m.getTileHeight(x-1, y) == currentHeight+1 || m.getTileHeight(x-1, y) <= currentHeight {
		coords = append(coords, coord{
			x: x - 1,
			y: y,
		})
	}
	if m.getTileHeight(x, y+1) == currentHeight+1 || m.getTileHeight(x, y+1) <= currentHeight {
		coords = append(coords, coord{
			x: x,
			y: y + 1,
		})
	}
	if m.getTileHeight(x, y-1) == currentHeight+1 || m.getTileHeight(x, y-1) <= currentHeight {
		coords = append(coords, coord{
			x: x,
			y: y - 1,
		})
	}
	return coords
}
