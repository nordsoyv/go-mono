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

type pathFragment struct {
	location coord
	cost     int
	fCost    int
	parent   *pathFragment
}

func (m Map) getLegalMoves(pos coord) []coord {
	coords := make([]coord, 0)
	currentHeight := m.getTileHeight(pos.x, pos.y)
	if m.getTileHeight(pos.x+1, pos.y) == currentHeight+1 || m.getTileHeight(pos.x+1, pos.y) <= currentHeight {
		coords = append(coords, coord{
			x: pos.x + 1,
			y: pos.y,
		})
	}
	if m.getTileHeight(pos.x-1, pos.y) == currentHeight+1 || m.getTileHeight(pos.x-1, pos.y) <= currentHeight {
		coords = append(coords, coord{
			x: pos.x - 1,
			y: pos.y,
		})
	}
	if m.getTileHeight(pos.x, pos.y+1) == currentHeight+1 || m.getTileHeight(pos.x, pos.y+1) <= currentHeight {
		coords = append(coords, coord{
			x: pos.x,
			y: pos.y + 1,
		})
	}
	if m.getTileHeight(pos.x, pos.y-1) == currentHeight+1 || m.getTileHeight(pos.x, pos.y-1) <= currentHeight {
		coords = append(coords, coord{
			x: pos.x,
			y: pos.y - 1,
		})
	}
	return coords
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m Map) costToGoal(tile coord) int {
	return Abs(tile.x-m.end.x) + Abs(tile.y-m.end.y)
}

func findLowestCostFragmentPosition(list []pathFragment) int {
	lowest := 100000
	foundPos := 0
	for pos, fragment := range list {
		//fragment := fragment
		if fragment.fCost < lowest {
			lowest = fragment.fCost
			foundPos = pos
		}
	}
	return foundPos
}

func remove(s []pathFragment, i int) []pathFragment {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (m Map) aStar() {

	// A* Search Algorithm
	//1.  Initialize the open list
	open := make([]pathFragment, 0)
	//2.  Initialize the closed list
	closed := make([]pathFragment, 0)
	//    put the starting node on the open
	//    list (you can leave its f at zero)
	//
	start := pathFragment{
		location: m.start,
		cost:     0,
		parent:   nil,
		fCost:    m.costToGoal(m.start),
	}
	open = append(open, start)
	//3.  while the open list is not empty
	for true {
		//    a) find the node with the least f on
		//       the open list, call it "q"
		qPos := findLowestCostFragmentPosition(open)
		//    b) pop q off the open list
		q := open[qPos]
		open = remove(open, qPos)
		//    c) generate q's 4 successors and set their
		//       parents to q
		//    d) for each successor
		legalMoves := m.getLegalMoves(q.location)
		for _, move := range legalMoves {
			if move == m.end {
				//        i) if successor is the goal, stop search

			}
			//        ii) else, compute both g and h for successor
			//          successor.g = q.g + distance between
			//                              successor and q
			//          successor.h = distance from goal to
			//          successor (This can be done using many
			//          ways, we will discuss three heuristics-
			//          Manhattan, Diagonal and Euclidean
			//          Heuristics)
			//
			//          successor.f = successor.g + successor.h

			nextFragment := pathFragment{
				location: move,
				cost:     q.cost + 1,
				fCost:    q.cost + 1 + m.costToGoal(move),
				parent:   &q,
			}

		}

	}
	//
	//
	//

	//
	//
	//        iii) if a node with the same position as
	//            successor is in the OPEN list which has a
	//           lower f than successor, skip this successor
	//
	//        iV) if a node with the same position as
	//            successor  is in the CLOSED list which has
	//            a lower f than successor, skip this successor
	//            otherwise, add  the node to the open list
	//     end (for loop)
	//
	//    e) push q on the closed list
	//    end (while loop)

}
