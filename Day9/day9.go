package day9

import (
	"math"
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	instructions := createInstructions(lines)
	return puzzle1(instructions), puzzle2(instructions)
}

func createInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0)

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])

		instructions = append(instructions, Instruction{
			direction: direction,
			distance:  distance,
		})
	}

	return instructions
}

func puzzle1(instructions []Instruction) int {
	rope := newRope(2)

	return solve(rope, instructions)
}

func puzzle2(instructions []Instruction) int {
	rope := newRope(10)

	return solve(rope, instructions)
}

func solve(rope Rope, instructions []Instruction) int {
	pointsVisitedByTail := make([]common.Point, 0)

	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			rope.knots[0] = rope.knots[0].Translate(instruction.direction)

			for j := 0; j < len(rope.knots)-1; j++ {
				head := rope.knots[j]
				tail := rope.knots[j+1]

				if tooFarAway(head, tail) {
					var candidates []common.Point
					if head.X == tail.X || head.Y == tail.Y {
						candidates = []common.Point{
							{X: tail.X + 1, Y: tail.Y},
							{X: tail.X - 1, Y: tail.Y},
							{X: tail.X, Y: tail.Y + 1},
							{X: tail.X, Y: tail.Y - 1},
						}
					} else {
						candidates = []common.Point{
							{X: tail.X + 1, Y: tail.Y + 1},
							{X: tail.X + 1, Y: tail.Y - 1},
							{X: tail.X - 1, Y: tail.Y + 1},
							{X: tail.X - 1, Y: tail.Y - 1},
						}
					}

					dist := head.GetManhattanDistance(candidates[0])
					closest := candidates[0]
					for k := 0; k < len(candidates); k++ {
						if head.GetManhattanDistance(candidates[k]) <= dist {
							dist = head.GetManhattanDistance(candidates[k])
							closest = candidates[k]
						}
					}

					rope.knots[j+1] = closest

				} else {
					continue
				}
			}
			pointsVisitedByTail = addIfNew(pointsVisitedByTail, rope.knots[len(rope.knots)-1])
		}
	}

	return len(pointsVisitedByTail)
}

type Rope struct {
	knots []common.Point
}

func newRope(length int) Rope {
	r := Rope{}
	r.knots = make([]common.Point, length)

	for i := 0; i < length; i++ {
		r.knots[i] = common.Point{X: 0, Y: 0}
	}

	return r
}

type Instruction struct {
	direction string
	distance  int
}

func addIfNew(ps []common.Point, newPoint common.Point) []common.Point {
	seen := false
	for _, point := range ps {
		seen = seen || point.Equals(newPoint)
	}

	if !seen {
		ps = append(ps, newPoint)
	}

	return ps
}

func tooFarAway(p1 common.Point, p2 common.Point) bool {
	xDist := int(math.Abs(float64(p1.X - p2.X)))
	yDist := int(math.Abs(float64(p1.Y - p2.Y)))

	return xDist > 1 || yDist > 1
}
