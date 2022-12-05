package days

import (
	"aoc_2022/common"
	"fmt"
	"strings"
)

// A == Rock
// B == Paper
// C == Scissors
// X == Rock
// Y == Paper
// Z == Scissors

func Day02A() {
	println("Day 2 A")
	score := testStrategyA("./res/day2a.txt")
	fmt.Println(score)
}

func Day02B() {
	println("Day 2 B")
	score := testStrategyB("./res/day2a.txt")
	fmt.Println(score)
}

func Day02TestA() {
	println("Day 2 test A")
	score := testStrategyA("./res/day2test.txt")
	fmt.Println(score)
}

func Day02TestB() {
	println("Day 2 test B")
	score := testStrategyB("./res/day2test.txt")
	fmt.Println("Scoure should be 12. Is actually", score)
}

func testStrategyA(path string) int {
	lines := common.ReadFileToLines(path)
	score := 0
	for _, line := range lines {
		score += evalMoveA(line)
	}
	return score
}
func testStrategyB(path string) int {
	lines := common.ReadFileToLines(path)
	score := 0
	for _, line := range lines {
		score += evalMoveB(line)
	}
	return score
}

func evalMoveA(move string) int {
	parts := strings.Split(move, " ")
	opponentMove := parts[0]
	myMove := parts[1]
	shapeScore := 0
	outcomeScore := 0
	if myMove == "X" {
		shapeScore = 1
	} else if myMove == "Y" {
		shapeScore = 2
	} else {
		shapeScore = 3
	}
	if opponentMove == "A" {
		if myMove == "X" { // draw
			outcomeScore = 3
		} else if myMove == "Y" { // win
			outcomeScore = 6
		} else {
			outcomeScore = 0
		}
	} else if opponentMove == "B" {
		if myMove == "X" {
			outcomeScore = 0
		} else if myMove == "Y" {
			outcomeScore = 3
		} else {
			outcomeScore = 6
		}
	} else { // C
		if myMove == "X" {
			outcomeScore = 6
		} else if myMove == "Y" {
			outcomeScore = 0
		} else {
			outcomeScore = 3
		}
	}
	return shapeScore + outcomeScore
}

func evalMoveB(move string) int {
	// X == should lose
	// Y ==should draw
	// Z == Should win
	// win = 6
	// draw = 3
	// lose = 0
	// Rock = 1
	// paper = 2
	// scissors = 3

	parts := strings.Split(move, " ")
	opponentMove := parts[0]
	myMove := parts[1]
	shapeScore := 0
	outcomeScore := 0
	if opponentMove == "A" { // rook
		if myMove == "X" { // lose
			outcomeScore = 0
			shapeScore = 3
		} else if myMove == "Y" { // draw
			outcomeScore = 3
			shapeScore = 1
		} else { // win
			outcomeScore = 6
			shapeScore = 2
		}
	} else if opponentMove == "B" { // paper
		if myMove == "X" { // lose
			outcomeScore = 0
			shapeScore = 1
		} else if myMove == "Y" { // draw
			outcomeScore = 3
			shapeScore = 2
		} else { // win
			outcomeScore = 6
			shapeScore = 3
		}
	} else { // C // scissors
		if myMove == "X" { // lose
			outcomeScore = 0
			shapeScore = 2
		} else if myMove == "Y" { // draw
			outcomeScore = 3
			shapeScore = 3
		} else { // win
			outcomeScore = 6
			shapeScore = 1
		}
	}
	return shapeScore + outcomeScore
}
