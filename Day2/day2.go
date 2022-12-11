package day2

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
	throws := make(map[byte]string)

	throws['A'] = "Rock"
	throws['X'] = "Rock"

	throws['B'] = "Paper"
	throws['Y'] = "Paper"

	throws['C'] = "Scissors"
	throws['Z'] = "Scissors"

	scoreForChosenThrow := make(map[string]int)

	scoreForChosenThrow["Rock"] = 1
	scoreForChosenThrow["Paper"] = 2
	scoreForChosenThrow["Scissors"] = 3

	var lines []string = common.Stuff(input)
	var score int = 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		them := throws[[]byte(splitLine[0])[0]]
		me := throws[[]byte(splitLine[1])[0]]

		score += scoreForChosenThrow[me]

		if them == me {
			score += 3
		} else if (them == "Rock" && me == "Paper") || (them == "Paper" && me == "Scissors") || (them == "Scissors" && me == "Rock") {
			score += 6
		}
	}
	return score
}

func puzzle2() int {
	throws := make(map[byte]string)

	throws['A'] = "Rock"
	throws['B'] = "Paper"
	throws['C'] = "Scissors"

	scoreForOutcome := make(map[byte]int)

	scoreForOutcome['X'] = 0
	scoreForOutcome['Y'] = 3
	scoreForOutcome['Z'] = 6

	var lines []string = common.Stuff(input)
	var score int = 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		them := throws[[]byte(splitLine[0])[0]]
		outcome := []byte(splitLine[1])[0]

		score += scoreForOutcome[outcome]

		var me string

		if outcome == 'Y' {
			me = them
		} else {
			switch them {
			case "Rock":
				if outcome == 'X' {
					me = "Scissors"
				} else {
					me = "Paper"
				}
			case "Paper":
				if outcome == 'X' {
					me = "Rock"
				} else {
					me = "Scissors"
				}
			case "Scissors":
				if outcome == 'X' {
					me = "Paper"
				} else {
					me = "Rock"
				}
			}
		}

		switch me {
		case "Rock":
			score += 1
		case "Paper":
			score += 2
		case "Scissors":
			score += 3
		}
	}

	return score
}
