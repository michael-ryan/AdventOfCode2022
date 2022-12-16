package day11

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
)

const (
	maxInt = math.MaxInt64
)

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func createMonkeyArray(lines []string) []*monkey {
	unparsedMonkeys := make([][]string, 0)
	unparsedMonkey := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			unparsedMonkeys = append(unparsedMonkeys, unparsedMonkey)
			unparsedMonkey = make([]string, 0)
		} else {
			unparsedMonkey = append(unparsedMonkey, line)
		}
	}

	unparsedMonkeys = append(unparsedMonkeys, unparsedMonkey)

	monkeys := make([]*monkey, 0)

	for _, monkey := range unparsedMonkeys {
		monkeys = append(monkeys, parseMonkey(monkey))
	}

	return monkeys
}

func puzzle1(lines []string) int {
	monkeys := createMonkeyArray(lines)
	for i := 0; i < 20; i++ {
		monkeys = simulateRoundP1(monkeys)
	}

	return calculateMonkeyBusiness(monkeys)
}

func simulateRoundP1(monkeys []*monkey) []*monkey {
	for _, monkey := range monkeys {
		for i := range monkey.items {
			monkey.inspectionsDone++
			monkey.items[i] = monkey.operation(monkey.items[i])
			monkey.items[i] = int(math.Floor(float64(monkey.items[i]) / 3))
			if monkey.test(monkey.items[i]) {
				monkeys[monkey.trueTarget].items = append(monkeys[monkey.trueTarget].items, monkey.items[i])
			} else {
				monkeys[monkey.falseTarget].items = append(monkeys[monkey.falseTarget].items, monkey.items[i])
			}
		}
		monkey.items = make([]int, 0)
	}

	return monkeys
}

func puzzle2(lines []string) int {
	monkeys := createMonkeyArray(lines)

	max := 1

	for _, monkey := range monkeys {
		max *= monkey.divisibilityBase
	}

	for i := 0; i < 10000; i++ {
		monkeys = simulateRoundP2(monkeys, max)
	}

	return calculateMonkeyBusiness(monkeys)
}

func simulateRoundP2(monkeys []*monkey, productOfMonkeyDivBases int) []*monkey {
	for _, monkey := range monkeys {
		for i := range monkey.items {
			monkey.inspectionsDone++
			monkey.items[i] = monkey.operation(monkey.items[i])
			monkey.items[i] = monkey.items[i] % productOfMonkeyDivBases
			if monkey.test(monkey.items[i]) {
				monkeys[monkey.trueTarget].items = append(monkeys[monkey.trueTarget].items, monkey.items[i])
			} else {
				monkeys[monkey.falseTarget].items = append(monkeys[monkey.falseTarget].items, monkey.items[i])
			}
		}
		monkey.items = make([]int, 0)
	}

	return monkeys
}

func check(i int) {
	if i == maxInt {
		panic(i)
	}
}

func calculateMonkeyBusiness(monkeys []*monkey) int {
	monkeysCopy := monkeys[:]

	sort.Slice(monkeysCopy, func(i, j int) bool {
		return monkeysCopy[i].inspectionsDone < monkeysCopy[j].inspectionsDone
	})

	return monkeysCopy[len(monkeysCopy)-1].inspectionsDone * monkeysCopy[len(monkeysCopy)-2].inspectionsDone
}

type monkey struct {
	inspectionsDone  int
	items            []int
	operation        func(old int) int
	test             func(worry int) bool
	trueTarget       int
	falseTarget      int
	divisibilityBase int
}

func parseMonkey(unparsedMonkey []string) *monkey {
	monkey := monkey{}
	monkey.inspectionsDone = 0
	items := strings.Split(strings.Split(unparsedMonkey[1], ": ")[1], ", ")
	for _, item := range items {
		itemAsInt, _ := strconv.Atoi(item)
		monkey.items = append(monkey.items, itemAsInt)
	}

	sum := strings.Split(unparsedMonkey[2], " = ")[1]

	operator := sum[4]

	operand := sum[6:]

	if operand == "old" {
		monkey.operation = func(old int) int { return old * old }
	} else {
		switch operand, _ := strconv.Atoi(operand); operator {
		case '*':
			monkey.operation = func(old int) int { return old * operand }
		case '+':
			monkey.operation = func(old int) int { return old + operand }
		}
	}

	monkey.divisibilityBase, _ = strconv.Atoi(strings.Split(unparsedMonkey[3], " by ")[1])

	monkey.test = func(worry int) bool { return worry%monkey.divisibilityBase == 0 }

	monkey.trueTarget, _ = strconv.Atoi(strings.Split(unparsedMonkey[4], " monkey ")[1])
	monkey.falseTarget, _ = strconv.Atoi(strings.Split(unparsedMonkey[5], " monkey ")[1])

	return &monkey
}
