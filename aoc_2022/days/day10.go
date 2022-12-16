package days

import (
	"aoc_2022/common"
	"fmt"
	"strings"
)

func Day10A() {
	println("Day 10 A")
	task10a("./res/day10.txt", 13820)
}

func Day10B() {
	println("Day 10 B")
	task10b("./res/day10.txt", -1)
}

func Day10TestA() {
	println("Day 10 test A")
	task10a("./res/day10test.txt", 13140)

}

func Day10TestB() {
	println("Day 10 test B")
	task10b("./res/day10test.txt", -1)
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
	//fmt.Println(machine.log)
	valueAt20 := machine.GetCycleValue(20)
	valueAt60 := machine.GetCycleValue(60)
	valueAt100 := machine.GetCycleValue(100)
	valueAt140 := machine.GetCycleValue(140)
	valueAt180 := machine.GetCycleValue(180)
	valueAt220 := machine.GetCycleValue(220)
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

func task10b(path string, _ int) {
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
	machine.DrawCrt()

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

func (m Machine) GetCycleValue(cycle int) int {
	offset := -1
	if (cycle + offset) < 0 {
		return 0
	}
	return m.log[cycle+offset]
}

func (m Machine) DrawCrt() {
	crtLine := ""
	for i := 0; i < len(m.log); i++ {
		if i%40 == 0 {
			fmt.Println(crtLine)
			crtLine = ""
		}
		logValue := m.GetCycleValue(i)
		if pixelHit(i%40, logValue) {
			crtLine += "#"
		} else {
			crtLine += " "
		}

	}
	fmt.Println(crtLine)

}

func pixelHit(cycle, pos int) bool {
	if cycle == pos || cycle == pos+1 || cycle == pos-1 {
		return true
	}
	return false
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
