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
}

func Day09TestA() {
	println("Day 09 test A")
	task9A("./res/day9test.txt", 13)
}

func Day09TestB() {
	println("Day 09 test B")
}

func task9A(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	commands := make([]Command, 0)
	for _, line := range lines {
		command := parseCommand(line)
		commands = append(commands, command)
	}
	tail := Position{
		x: 0,
		y: 0,
	}
	head := Position{
		x: 0,
		y: 0,
	}
	positions := map[Position]bool{}
	positions[tail] = true

	for _, command := range commands {
		dir := command.dir
		for i := 0; i < command.amount; i++ {
			switch dir {
			case "R":
				head.x++
			case "L":
				head.x--
			case "D":
				head.y--
			case "U":
				head.y++
			}
			//fmt.Println(head)
			dist := getDistance(head, tail)
			if dist > 1 {
				if head.x == tail.x { // need to move vertical
					if tail.y < head.y {
						tail.y++
					} else {
						tail.y--
					}
				} else if head.y == tail.y { // need to move horizontal
					if tail.x < head.x {
						tail.x++
					} else {
						tail.x--
					}
				} else { // need to move diagonal
					if tail.x < head.x {
						tail.x++
					} else {
						tail.x--
					}
					if tail.y < head.y {
						tail.y++
					} else {
						tail.y--
					}

				}
			}
			positions[tail] = true
		}
	}

	result = len(positions)
	fmt.Println("Number of positions occupied by tail:", result, "Should be ", expected)
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
