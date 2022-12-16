package day10

import (
	"fmt"
	"testing"
)

var samplePath string = "../resources/samples/day10.txt"

func TestPuzzle1(t *testing.T) {
	actual, _ := Answers(samplePath)

	var expected int = 13140

	message := fmt.Sprintf("Expected %v, got %v\n", expected, actual)

	if expected == actual {
		t.Logf(message)
	} else {
		t.Errorf(message)
	}
}

func TestPuzzle2(t *testing.T) {
	_, actual := Answers(samplePath)

	var expected string = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

	message := fmt.Sprintf("Expected %v, got %v\n", expected, actual)

	if expected == actual {
		t.Logf(message)
	} else {
		t.Errorf(message)
	}
}
