package day6

import (
	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (int, int) {
	line := common.ReadFileAsStringArray(inputPath)[0]
	return puzzle1(line), puzzle2(line)
}

func puzzle1(line string) int {
	for i := 0; i < len(line)-4; i++ {
		substring := line[i : i+4]
		if unique(substring) {
			return i + 4
		}
	}
	panic("No solution found")
}

func puzzle2(line string) int {
	return 0
}

// from https://codereview.stackexchange.com/a/251210/49121
func unique(s string) bool {
	m := make(map[rune]bool)
	for _, i := range s {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
