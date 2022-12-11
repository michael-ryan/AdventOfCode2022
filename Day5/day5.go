package day5

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (string, string) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func puzzle1(lines []string) string {
	state, instructions := parseInput(lines)

	doInstructions(state, instructions, func(state []common.Stack[byte], instruction Instruction) {
		for i := 0; i < instruction.quantity; i++ {
			state[instruction.destination].Push(state[instruction.source].Pop())
		}
	})

	return extractAnswer(state)
}

func puzzle2(lines []string) string {
	state, instructions := parseInput(lines)

	doInstructions(state, instructions, func(state []common.Stack[byte], instruction Instruction) {
		tempStack := common.NewStack[byte]()
		for i := 0; i < instruction.quantity; i++ {
			tempStack.Push(state[instruction.source].Pop())
		}
		for tempStack.Size() > 0 {
			state[instruction.destination].Push(tempStack.Pop())
		}
	})

	return extractAnswer(state)
}

func doInstructions(state []common.Stack[byte], instructions []Instruction, doInstruction func(state []common.Stack[byte], instruction Instruction)) {
	for _, instruction := range instructions {
		doInstruction(state, instruction)
	}
}

func parseInput(lines []string) ([]common.Stack[byte], []Instruction) {
	var blankLine int = 0

	for number, line := range lines {
		if line == "" {
			blankLine = number
		}
	}

	state := parseInitialState(lines[:blankLine])
	unparsedInstructions := lines[blankLine+1:]

	instructions := make([]Instruction, len(unparsedInstructions))

	for i := 0; i < len(unparsedInstructions); i++ {
		instructions = append(instructions, parseInstruction(unparsedInstructions[i]))
	}

	return state, instructions
}

func extractAnswer(state []common.Stack[byte]) string {
	solution := ""

	for _, stack := range state {
		solution += string(stack.Peek())
	}

	return solution
}

func getNumberOfStacks(stackNumbersLine string) int {
	stackNumbers := strings.Split(stackNumbersLine, "   ")
	stackNumbers[0] = stackNumbers[0][1:]
	stackNumbers[len(stackNumbers)-1] = stackNumbers[len(stackNumbers)-1][:len(stackNumbers[len(stackNumbers)-1])-1]
	numberOfStacks, _ := strconv.Atoi(stackNumbers[len(stackNumbers)-1])
	return numberOfStacks
}

func parseInitialState(data []string) []common.Stack[byte] {
	numberOfStacks := getNumberOfStacks(data[len(data)-1])

	stacks := make([]common.Stack[byte], numberOfStacks)

	for i := len(data) - 2; i >= 0; i-- {
		line := data[i]
		for j := 0; j < numberOfStacks; j++ {
			if line[1+(4*j)] != ' ' {
				stacks[j].Push(line[(4*j)+1])
			}
		}
	}

	return stacks
}

func parseInstruction(prose string) Instruction {
	halves := strings.Split(prose, " from ")
	locations := strings.Split(halves[1], " ")

	quantity, _ := strconv.Atoi(strings.Split(halves[0], " ")[1])
	source, _ := strconv.Atoi(locations[0])
	destination, _ := strconv.Atoi(locations[2])

	i := Instruction{
		quantity:    quantity,
		source:      source - 1,
		destination: destination - 1,
	}

	return i
}

type Instruction struct {
	quantity    int
	source      int
	destination int
}
