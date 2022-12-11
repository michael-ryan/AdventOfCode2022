package day4

import (
	"sort"
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
	"golang.org/x/exp/slices"
)

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func puzzle1(lines []string) int {
	eclipseOccurences := 0
	for _, line := range lines {
		elf1, elf2 := parseLine(line)

		elf1Slice := elf1.AsSlice()
		sort.Ints(elf1Slice)

		elf2Slice := elf2.AsSlice()
		sort.Ints(elf2Slice)

		intersection := elf1.Intersect(elf2).AsSlice()
		sort.Ints(intersection)

		if slices.Equal(intersection, elf1Slice) || slices.Equal(intersection, elf2Slice) {
			eclipseOccurences++
		}
	}

	return eclipseOccurences
}

func puzzle2(lines []string) int {
	overlapOccurences := 0
	for _, line := range lines {
		elf1, elf2 := parseLine(line)

		elf1Slice := elf1.AsSlice()
		elf2Slice := elf2.AsSlice()

		for _, elf1Zone := range elf1Slice {
			if contains(elf2Slice, elf1Zone) {
				overlapOccurences++
				break
			}
		}
	}

	return overlapOccurences
}

func contains(is []int, target int) bool {
	for _, v := range is {
		if v == target {
			return true
		}
	}

	return false
}

func parseLine(line string) (*common.Set[int], *common.Set[int]) {
	splitLine := strings.Split(line, ",")
	elf1 := parseElf(splitLine[0])
	elf2 := parseElf(splitLine[1])

	return elf1, elf2
}

func parseElf(elf string) *common.Set[int] {
	splitElf := strings.Split(elf, "-")
	from, _ := strconv.Atoi(splitElf[0])
	to, _ := strconv.Atoi(splitElf[1])

	s := common.NewSet[int]()

	for i := from; i <= to; i++ {
		s.Add(i)
	}

	return s
}
