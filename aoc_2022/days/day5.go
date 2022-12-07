package days

import (
	"aoc_2022/common"
	"fmt"
	"regexp"
)

func Day05A() {
	println("Day 05 A")
	s, moves := parseFile("./res/day5.txt")
	for _, move := range moves {
		s.applyMove9000(move)
		//fmt.Println(s)
	}
	fmt.Println("Answer is", s.getTopCrates(), "Should be HBTMTBSDC")
}

func Day05B() {
	println("Day 05 B")
	s, moves := parseFile("./res/day5.txt")
	for _, move := range moves {
		s.applyMove9001(move)
		//fmt.Println(s)
	}
	fmt.Println("Answer is", s.getTopCrates(), "Should be PQTJRSHWS")
}

func Day05TestA() {
	println("Day 05 test A")
	s, moves := parseFile("./res/day5test.txt")
	for _, move := range moves {
		s.applyMove9000(move)
		fmt.Println(s)
	}
	fmt.Println("Answer is", s.getTopCrates(), "Should be CMZ")
}

func Day05TestB() {
	println("Day 05 test B")
	s, moves := parseFile("./res/day5test.txt")
	for _, move := range moves {
		s.applyMove9001(move)
		fmt.Println(s)
	}
	fmt.Println("Answer is", s.getTopCrates(), "Should be MCD")
}

type Stacks struct {
	stacks [9][]string
}

type Move struct {
	amount int
	from   int
	to     int
}

func (m Move) String() string {
	return fmt.Sprintf("move %v from %v to %v\n", m.amount, m.from, m.to)
}

func createStacks() Stacks {
	s := Stacks{stacks: [9][]string{}}
	return s
}

func createMove(amount, from, to string) Move {
	return Move{
		amount: common.StringToInt(amount),
		from:   common.StringToInt(from),
		to:     common.StringToInt(to),
	}
}

func (s *Stacks) addCrate(index int, value string) {
	current := s.stacks[index]
	if current == nil {
		current = []string{}
	}
	s.stacks[index] = append(s.stacks[index], value)
}

func (s Stacks) String() string {
	numRows := 0
	for _, stack := range s.stacks {
		length := len(stack)
		if length > numRows {
			numRows = length
		}
	}
	allRows := ""
	for i := 0; i < numRows; i++ {
		currentRow := ""

		for _, stack := range s.stacks {
			stackLen := len(stack)
			if numRows-i > stackLen {
				currentRow += "    "
			} else {
				currentRow += fmt.Sprintf("[%v] ", stack[i-(numRows-stackLen)])
			}
		}
		allRows += currentRow + "\n"
	}
	allRows += " 1   2   3   4   5   6   7   8   9\n"
	return allRows
}

func (s *Stacks) applyMove9000(move Move) {
	newSlice, toMove := common.RemoveFromFrontOfSlice(s.stacks[move.from-1], move.amount)
	s.stacks[move.from-1] = newSlice
	s.stacks[move.to-1] = common.AddToFrontOfSlice(s.stacks[move.to-1], toMove)
}

func (s *Stacks) applyMove9001(move Move) {
	newSlice, toMove := common.RemoveFromFrontOfSlice(s.stacks[move.from-1], move.amount)
	s.stacks[move.from-1] = newSlice
	s.stacks[move.to-1] = common.AddStackToFrontOfSlice(s.stacks[move.to-1], toMove)
}

func (s Stacks) getTopCrates() string {
	result := ""
	for _, stack := range s.stacks {
		if len(stack) > 0 {
			result += stack[0]
		} else {
			result += " "
		}

	}
	return result
}

func parseFile(path string) (Stacks, []Move) {
	lines := common.ReadFileToLines(path)
	stack := createStacks()
	moves := []Move{}
	parseStack := true
	for _, line := range lines {
		if parseStack {
			lineLength := len(line)
			if lineLength == 0 {
				parseStack = false
				continue
			}
			firstLetter := string(line[1])
			if firstLetter == "1" {
				continue
			}
			for i := 0; i < (lineLength/4)+1; i++ {
				index := (i * 3) + i + 1
				value := string(line[index])
				if value != " " {
					stack.addCrate(i, value)
				}
			}
		} else {
			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllString(line, -1)
			moves = append(moves, createMove(matches[0], matches[1], matches[2]))
		}
	}
	return stack, moves
}
