package days

import (
	"aoc_2022/common"
	"fmt"
	"math"
	"strings"
)

func Day09A() {
	println("Day 09 A")
	task9A("./res/day9.txt", 6098)
}

func Day09B() {
	println("Day 09 B")
	task9B("./res/day9.txt", 0)
}

func Day09TestA() {
	println("Day 09 test A")
	task9A("./res/day9test.txt", 13)
}

func Day09TestB() {
	println("Day 09 test B")
	task9B("./res/day9test.txt", 1)
	task9B("./res/day9test2.txt", 36)
}

func task9A(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	commands := make([]Command, 0)
	for _, line := range lines {
		command := parseCommand(line)
		commands = append(commands, command)
	}
	result = moveRope(2, commands)
	fmt.Println("Number of positions occupied by tail:", result, "Should be ", expected)
}
func task9B(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	commands := make([]Command, 0)
	for _, line := range lines {
		command := parseCommand(line)
		commands = append(commands, command)
	}
	result = moveRope(10, commands)
	fmt.Println("Number of positions occupied by tail:", result, "Should be ", expected)
}

func moveRope(ropeLength int, commands []Command) int {
	rope := make([]Position, ropeLength)
	positions := map[Position]bool{}
	positions[rope[len(rope)-1]] = true
	for _, command := range commands {
		dir := command.dir
		for i := 0; i < command.amount; i++ {
			switch dir {
			case "R":
				rope[0].x++
			case "L":
				rope[0].x--
			case "D":
				rope[0].y--
			case "U":
				rope[0].y++
			}
			for i := 1; i < len(rope); i++ {
				rope[i].MoveToward(rope[i-1])
			}
			positions[rope[len(rope)-1]] = true
		}
	}
	return len(positions)
}

type Position struct {
	x, y int
}

type Command struct {
	dir    string
	amount int
}

func (c Command) String() string {
	return fmt.Sprint(c.dir, " ", c.amount)
}

func parseCommand(command string) Command {
	parts := strings.Split(command, " ")
	amount := common.StringToInt(parts[1])
	return Command{
		dir:    parts[0],
		amount: amount,
	}
}

func getDistance(a, b Position) int {
	return int(math.Max(math.Abs(float64(a.x-b.x)), math.Abs(float64(a.y-b.y))))
}

func (p *Position) MoveToward(target Position) {
	dist := getDistance(target, *p)
	if dist < 2 {
		return
	}
	if target.x == p.x { // need to move vertical
		if p.y < target.y {
			p.y++
		} else {
			p.y--
		}
	} else if target.y == p.y { // need to move horizontal
		if p.x < target.x {
			p.x++
		} else {
			p.x--
		}
	} else { // need to move diagonal
		if p.x < target.x {
			p.x++
		} else {
			p.x--
		}
		if p.y < target.y {
			p.y++
		} else {
			p.y--
		}

	}
}
