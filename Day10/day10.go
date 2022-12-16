package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (int, string) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func puzzle1(lines []string) int {
	accumulatorHistory := generateAccumulatorHistory(lines)

	return computeAnswerP1(accumulatorHistory)
}

func computeAnswerP1(numbers []int) int {
	return (20 * numbers[19]) + (60 * numbers[59]) + (100 * numbers[99]) + (140 * numbers[139]) + (180 * numbers[179]) + (220 * numbers[219])
}

func puzzle2(lines []string) string {
	accumulatorHistory := generateAccumulatorHistory(lines)

	return generatePrintedString(accumulatorHistory)
}

func generatePrintedString(accumulatorHistory []int) string {
	printedString := ""

	for i := 0; i < 240; i += 40 {
		printedString += generateLine(accumulatorHistory[i:i+40]) + "\n"
	}

	return printedString[:len(printedString)-1]
}

func generateLine(fortyCyclesOfAccumulator []int) string {
	if len(fortyCyclesOfAccumulator) != 40 {
		panic(fmt.Sprintf("Expected 40 cycles, got %v", len(fortyCyclesOfAccumulator)))
	}

	line := ""

	for i := 0; i < 40; i++ {
		currentValue := fortyCyclesOfAccumulator[i]
		if currentValue <= i+1 && currentValue >= i-1 {
			line += "#"
		} else {
			line += "."
		}
	}

	return line
}

func generateAccumulatorHistory(lines []string) []int {
	accumulatorHistory := make([]int, 0)
	accumulator := 1
	for _, line := range lines {
		if line == "noop" {
			accumulatorHistory = append(accumulatorHistory, accumulator)
		} else {
			parts := strings.Split(line, " ")
			k, _ := strconv.Atoi(parts[1])

			accumulatorHistory = append(accumulatorHistory, accumulator)
			accumulatorHistory = append(accumulatorHistory, accumulator)

			accumulator += k
		}
	}

	return accumulatorHistory
}
