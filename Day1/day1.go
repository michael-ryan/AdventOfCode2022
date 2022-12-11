package day1

import (
	_ "embed"
	"sort"
	"strconv"

	"github.com/michael-ryan/AoC22/common"
)

var (
	//go:embed input.txt
	input string
)

func Answers() (int, int) {
	return puzzle1(), puzzle2()
}

func calculateTotalCaloriesByElf() []int {
	var lines []string = common.Stuff(input)
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

func puzzle1() int {
	sums := calculateTotalCaloriesByElf()

	return sums[len(sums)-1]
}

func puzzle2() int {
	sums := calculateTotalCaloriesByElf()

	return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}
