package days

import (
	"aoc_2022/common"
	"testing"
)

func TestDay1(t *testing.T) {
	lines := common.ReadFileToLines("../res/day1test.txt")
	if len(lines) != 14 {
		t.Fatalf("Wrong number of lines. Got %v, exptede 15", len(lines))
	}
	for _, line := range lines {
		t.Log(line)
	}
}
