package day1

import (
	"sort"
	"strconv"

	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func calculateTotalCaloriesByElf(lines []string) []int {
	var sums []int = make([]int, 0)
	var total int
	for _, line := range lines {
		if line == "" {
			sums = append(sums, total)
			total = 0
		} else {
			calories, _ := strconv.Atoi(line)
			total += calories
		}
	}

	sums = append(sums, total)

	sort.Ints(sums)

	return sums
}

func puzzle1(lines []string) int {
	sums := calculateTotalCaloriesByElf(lines)

	return sums[len(sums)-1]
}

func puzzle2(lines []string) int {
	sums := calculateTotalCaloriesByElf(lines)

	return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}
