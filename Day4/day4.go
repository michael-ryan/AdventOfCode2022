package day4

import (
	"sort"
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
	"golang.org/x/exp/slices"
)

func Puzzle1() int {
	lines := common.ReadFileAsStringArray("Day4/input.txt")

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

func Puzzle2() int {
	lines := common.ReadFileAsStringArray("Day4/input.txt")

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