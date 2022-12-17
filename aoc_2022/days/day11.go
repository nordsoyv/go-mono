package days

import (
	"aoc_2022/common"
	"fmt"
	"sort"
	"strings"
)

func Day11A() {
	println("Day 11 A")
	task11a("./res/day11.txt", 58786)
}

func Day11B() {
	println("Day 11 B")
	task11b("./res/day11.txt", -1)
}

func Day11TestA() {
	println("Day 11 test A")
	task11a("./res/day11test.txt", 10605)
}

func Day11TestB() {
	println("Day 11 test B")
	task11b("./res/day11.txt", -1)
}

func task11a(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	monkeys := parseMonkeys(lines)
	fmt.Println(monkeys)
	for round := 0; round < 20; round++ {
		fmt.Println("ROUND", round)
		for i, monkey := range monkeys {
			monkey := monkey
			for _, item := range monkey.items {
				monkeys[i].inspectAmount++
				worry := item
				if monkey.operation.operator == "+" {
					if monkey.operation.num == "old" {
						worry += item
					} else {
						value := common.StringToInt(monkey.operation.num)
						worry += value
					}
				} else if monkey.operation.operator == "*" {
					if monkey.operation.num == "old" {
						worry *= item
					} else {
						value := common.StringToInt(monkey.operation.num)
						worry *= value
					}
				} else {
					panic(fmt.Sprintf("Unknown opeartor %v", monkey.operation.operator))
				}
				//fmt.Printf("Old Worry lvl %v, new worry lvl %v\n", item, worry)
				worry /= 3
				//fmt.Printf("Old Worry lvl %v, new worry lvl %v\n", item, worry)
				if worry%monkey.test == 0 {
					//fmt.Printf("Throw to monkey %v\n", monkey.trueTarget)
					monkeys[monkey.trueTarget].addItem(worry)
				} else {
					//fmt.Printf("Throw to monkey %v\n", monkey.falseTarget)
					monkeys[monkey.falseTarget].addItem(worry)
				}
			}
			monkeys[i].items = make([]int, 0)
		}

		fmt.Println("The monkeys are holding these items:")
		for i, monkey := range monkeys {
			fmt.Printf("Monkey %v is holding: %v\n", i, monkey.items)
		}
		fmt.Println("ROUND DONE")
	}
	inspectAmounts := make([]int, 0)
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %v has inspected: %v\n", i, monkey.inspectAmount)
		inspectAmounts = append(inspectAmounts, monkey.inspectAmount)
	}
	sort.Ints(inspectAmounts)
	fmt.Println(inspectAmounts)
	val1 := inspectAmounts[len(inspectAmounts)-1]
	val2 := inspectAmounts[len(inspectAmounts)-2]
	result = val1 * val2
	fmt.Printf("Sum monkey business %v. Expected %v", result, expected)
}

func runRounds(numRounds, worryDivider int, monkeys []Monkey) {
	for round := 0; round < numRounds; round++ {
		fmt.Println("ROUND", round)
		for i, monkey := range monkeys {
			monkey := monkey
			for _, item := range monkey.items {
				monkeys[i].inspectAmount++
				worry := item
				if monkey.operation.operator == "+" {
					if monkey.operation.num == "old" {
						worry += item
					} else {
						value := common.StringToInt(monkey.operation.num)
						worry += value
					}
				} else if monkey.operation.operator == "*" {
					if monkey.operation.num == "old" {
						worry *= item
					} else {
						value := common.StringToInt(monkey.operation.num)
						worry *= value
					}
				} else {
					panic(fmt.Sprintf("Unknown opeartor %v", monkey.operation.operator))
				}
				//fmt.Printf("Old Worry lvl %v, new worry lvl %v\n", item, worry)
				worry /= worryDivider
				//fmt.Printf("Old Worry lvl %v, new worry lvl %v\n", item, worry)
				if worry%monkey.test == 0 {
					//fmt.Printf("Throw to monkey %v\n", monkey.trueTarget)
					monkeys[monkey.trueTarget].addItem(worry)
				} else {
					//fmt.Printf("Throw to monkey %v\n", monkey.falseTarget)
					monkeys[monkey.falseTarget].addItem(worry)
				}
			}
			monkeys[i].items = make([]int, 0)
		}

		fmt.Println("The monkeys are holding these items:")
		for i, monkey := range monkeys {
			fmt.Printf("Monkey %v is holding: %v\n", i, monkey.items)
		}
		fmt.Println("ROUND DONE")
	}
}

func task11b(path string, expected int) {
	result := 0
	fmt.Printf("Sum monkey business %v. Expected %v", result, expected)
}

func parseMonkeys(lines []string) []Monkey {
	monkey := createMonkey()
	monkies := make([]Monkey, 0)
	monkey.id = len(monkies)
	for _, line := range lines {
		if line == "" {
			monkies = append(monkies, monkey)
			monkey = createMonkey()
			monkey.id = len(monkies)
			fmt.Println("")
			continue
		}
		if strings.HasPrefix(line, "Monkey") {
			continue
		}
		if strings.HasPrefix(line, "  Starting items") {
			itemStr := line[18:]
			items := strings.Split(itemStr, ", ")
			for _, item := range items {
				monkey.addItem(common.StringToInt(item))
			}
		}
		if strings.HasPrefix(line, "  Operation") {
			operationStr := line[23:]
			items := strings.Split(operationStr, " ")
			monkey.operation = operation{
				operator: items[0],
				num:      items[1],
			}
		}
		if strings.HasPrefix(line, "  Test") {
			test := line[21:]
			monkey.test = common.StringToInt(test)
		}
		if strings.HasPrefix(line, "    If true") {
			ifTrue := line[29:]
			monkey.trueTarget = common.StringToInt(ifTrue)
		}
		if strings.HasPrefix(line, "    If false") {
			ifFalse := line[30:]
			monkey.falseTarget = common.StringToInt(ifFalse)
		}
	}
	monkies = append(monkies, monkey)
	return monkies
}

type Monkey struct {
	id            int
	items         []int
	test          int
	operation     operation
	trueTarget    int
	falseTarget   int
	inspectAmount int
}

type operation struct {
	operator string
	num      string
}

func createMonkey() Monkey {
	return Monkey{
		id:    0,
		items: make([]int, 0),
		test:  0,
		operation: operation{
			operator: "",
			num:      "",
		},
		trueTarget:    0,
		falseTarget:   0,
		inspectAmount: 0,
	}
}

func (m *Monkey) addItem(item int) {
	m.items = append(m.items, item)
}

func (m Monkey) String() string {
	res := ""
	res = fmt.Sprintf("Monkey %v:\n", m.id)
	res += fmt.Sprintf("  Items: %v\n", m.items)
	res += fmt.Sprintf("  Operation: new = old %v %v\n", m.operation.operator, m.operation.num)
	res += fmt.Sprintf("  Test: divisible by %v\n", m.test)
	res += fmt.Sprintf("      If true: throw to monkey %v\n", m.trueTarget)
	res += fmt.Sprintf("      If false: throw to monkey %v\n", m.falseTarget)
	return res
}
