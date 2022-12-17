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
	task11b("./res/day11.txt", 14952185856)
}

func Day11TestA() {
	println("Day 11 test A")
	task11a("./res/day11test.txt", 10605)
}

func Day11TestB() {
	println("Day 11 test B")
	task11b("./res/day11test.txt", 2713310158)
}

func task11a(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	monkeys := parseMonkeys(lines)
	fmt.Println(monkeys)
	runRounds(20, 3, monkeys)
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

func task11b(path string, expected int) {
	result := 0
	lines := common.ReadFileToLines(path)
	monkeys := parseMonkeys(lines)
	commonDivider := int64(1)
	for _, monkey := range monkeys {
		commonDivider *= monkey.test
	}
	fmt.Println(monkeys, commonDivider)
	runRounds(10000, commonDivider, monkeys)
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

func runRounds(numRounds int, worryDivider int64, monkeys []Monkey) {
	for round := 0; round < numRounds; round++ {
		for i, monkey := range monkeys {
			monkey := monkey
			for _, item := range monkey.items {
				monkeys[i].inspectAmount++
				worry := item
				if monkey.operation.operator == "+" {
					if monkey.operation.num == "old" {
						worry += item
					} else {
						value := common.StringToInt64(monkey.operation.num)
						worry += value
					}
				} else if monkey.operation.operator == "*" {
					if monkey.operation.num == "old" {
						worry *= item
					} else {
						value := common.StringToInt64(monkey.operation.num)
						worry *= value
					}
				} else {
					panic(fmt.Sprintf("Unknown opeartor %v", monkey.operation.operator))
				}
				if worryDivider == 3 {
					worry /= worryDivider
				} else {
					worry %= worryDivider
				}
				if worry%monkey.test == 0 {
					monkeys[monkey.trueTarget].addItem(worry)
				} else {
					monkeys[monkey.falseTarget].addItem(worry)
				}
			}
			monkeys[i].items = make([]int64, 0)
		}
	}
}

func printInspectAmount(round int, monkeys []Monkey) {
	fmt.Println("Holding after round", round)
	for i, monkey := range monkeys {
		fmt.Printf("  Monkey %v has inspected: %v\n", i, monkey.inspectAmount)
	}

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
				monkey.addItem(common.StringToInt64(item))
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
			monkey.test = common.StringToInt64(test)
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
	items         []int64
	test          int64
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
		items: make([]int64, 0),
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

func (m *Monkey) addItem(item int64) {
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
