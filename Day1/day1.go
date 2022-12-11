package day1

import (
	"sort"
	"strconv"

	"github.com/michael-ryan/AoC22/common"
)

func calculateTotalCaloriesByElf() []int {
	var lines []string = common.ReadFileAsStringArray("Day1/input.txt")
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

	sort.Ints(sums)

	return sums
}

func Puzzle1() int {
	sums := calculateTotalCaloriesByElf()

	return sums[len(sums)-1]
}

func Puzzle2() int {
	sums := calculateTotalCaloriesByElf()

	return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}
