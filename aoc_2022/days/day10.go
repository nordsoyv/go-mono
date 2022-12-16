package days

import (
	"aoc_2022/common"
	"fmt"
	"strings"
)

func Day10A() {
	println("Day 10 A")
	task10a("./res/day10.txt", -1)
}

func Day10B() {
	println("Day 10 B")
}

func Day10TestA() {
	println("Day 10 test A")
	task10a("./res/day10test.txt", 13140)

}

func Day10TestB() {
	println("Day 10 test B")
}

func task10a(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	instructions := make([]Instruction, 0)
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}
	machine := Machine{
		log:          make([]int, 0),
		instructions: instructions,
		cycle:        1,
		regX:         1,
	}
	machine.ExecuteProgram()
	offset := -2
	//fmt.Println(machine.log)
	valueAt20 := machine.log[20+offset]
	valueAt60 := machine.log[60+offset]
	valueAt100 := machine.log[100+offset]
	valueAt140 := machine.log[140+offset]
	valueAt180 := machine.log[180+offset]
	valueAt220 := machine.log[220+offset]
	fmt.Println(valueAt20, 21)
	fmt.Println(valueAt60, 19)
	fmt.Println(valueAt100, 18)
	fmt.Println(valueAt140, 21)
	fmt.Println(valueAt180, 16)
	fmt.Println(valueAt220, 18)
	result += 20 * valueAt20
	result += 60 * valueAt60
	result += 100 * valueAt100
	result += 140 * valueAt140
	result += 180 * valueAt180
	result += 220 * valueAt220
	fmt.Println("Signal strength sum is:", result, "Should be ", expected)
}

type Machine struct {
	log          []int
	instructions []Instruction
	cycle        int
	regX         int
}

type Instruction struct {
	command string
	operand int
}

func (c Instruction) String() string {
	if c.command == "noop" {
		return fmt.Sprint(c.command)
	}
	return fmt.Sprint(c.command, " ", c.operand)
}

func (m *Machine) ExecuteProgram() {
	for _, instruction := range m.instructions {
		if instruction.command == "noop" {
			// noop
			fmt.Println("\tNoop")
			m.EndCycle()
		} else if instruction.command == "addx" {
			fmt.Println("\tAddx", instruction.operand)
			m.EndCycle()
			m.regX += instruction.operand
			m.EndCycle()
		}
	}
}

func (m *Machine) EndCycle() {
	m.log = append(m.log, m.regX)
	m.cycle++
	fmt.Println(len(m.log), m.regX)
}

func parseInstruction(input string) Instruction {
	parts := strings.Split(input, " ")
	if parts[0] == "noop" {
		return Instruction{
			command: parts[0],
			operand: 0,
		}
	}
	return Instruction{
		command: parts[0],
		operand: common.StringToInt(parts[1]),
	}
}
