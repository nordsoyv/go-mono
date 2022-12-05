package days

import (
	"aoc_2022/common"
	"fmt"
	"sort"
	"strconv"
)

type Elf struct {
	calories []int
}

func (e Elf) availableFood() int {
	var sum = 0
	for _, calory := range e.calories {
		sum += calory
	}
	return sum
}

func Day01A() {
	println("Day 1 A")
	maxFood := findElfWithMostFood("./res/day1a.txt")
	fmt.Println(maxFood)
}

func Day01B() {
	println("Day 1 B")
	maxFood := findTop3ElvesFood("./res/day1a.txt")
	fmt.Println(maxFood)
}

func Day01TestA() {
	println("Day 1 test A")
	maxFood := findElfWithMostFood("./res/day1test.txt")
	fmt.Println(maxFood)
}

func Day01TestB() {
	println("Day 1 test B")
	maxFood := findTop3ElvesFood("./res/day1test.txt")
	fmt.Println(maxFood)
}

func parseElves(path string) []Elf {
	lines := common.ReadFileToLines(path)
	elves := []Elf{}
	currentElf := Elf{}
	for _, line := range lines {
		if len(line) == 0 {
			elves = append(elves, currentElf)
			currentElf = Elf{}
		}
		parse, _ := strconv.ParseInt(line, 10, 32)
		currentElf.calories = append(currentElf.calories, int(parse))
	}
	elves = append(elves, currentElf)
	return elves
}

func findElfWithMostFood(path string) int {
	elves := parseElves(path)
	maxFood := 0
	for _, elf := range elves {
		currentElfFood := elf.availableFood()
		if currentElfFood > maxFood {
			maxFood = currentElfFood
		}
	}
	return maxFood
}

func findTop3ElvesFood(path string) int {
	elves := parseElves(path)
	var foodAmounts = []int{}
	for _, elf := range elves {
		foodAmounts = append(foodAmounts, elf.availableFood())
	}
	sort.Ints(foodAmounts)
	length := len(foodAmounts)
	return foodAmounts[length-1] + foodAmounts[length-2] + foodAmounts[length-3]
}
