package day3

import (
	_ "embed"
	"strings"

	"github.com/michael-ryan/AoC22/common"
)

var (
	//go:embed input.txt
	input string
)

func Answers() (int, int) {
	return puzzle1(), puzzle2()
}

func puzzle1() int {
	sumOfPriorities := 0
	for _, line := range common.Stuff(input) {
		length := len(line)
		commonItem := findCommonByte(line[:length/2], line[length/2:])
		sumOfPriorities += getItemPriority(commonItem)
	}

	return sumOfPriorities
}

func puzzle2() int {
	sumOfPriorities := 0
	lines := common.Stuff(input)
	for elfNumber := 0; elfNumber <= len(lines)-3; elfNumber += 3 {
		bag1 := stringToSet(lines[elfNumber])
		bag2 := stringToSet(lines[elfNumber+1])
		bag3 := stringToSet(lines[elfNumber+2])

		commonByte := bag1.Intersect(bag2).Intersect(bag3)

		// digusting hack to extract the one item in the final set post-intersections
		for _, char := range lines[elfNumber] {
			if commonByte.Has(byte(char)) {
				sumOfPriorities += getItemPriority(byte(char))
				break
			}
		}
	}

	return sumOfPriorities
}

func stringToSet(s string) *common.Set[byte] {
	set := common.NewSet[byte]()
	for _, char := range s {
		set.Add(byte(char))
	}

	return set
}

func findCommonByte(a string, b string) byte {
	for _, char := range a {
		if strings.Contains(b, string(char)) {
			return byte(char)
		}
	}

	panic("No common byte")
}

func getItemPriority(item byte) int {
	if item >= 97 && item <= 122 {
		// lowercase
		return int(item - 'a' + 1)
	} else if item >= 65 && item <= 90 {
		// uppercase
		return int(item - 'A' + 27)
	} else {
		panic(item)
	}
}
