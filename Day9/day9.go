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
	pointsVisitedByTail := make([]Point, 0)

	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			rope.knots[0] = rope.knots[0].translate(instruction.direction)

			for j := 0; j < len(rope.knots)-1; j++ {
				head := rope.knots[j]
				tail := rope.knots[j+1]

				if head.tooFarAway(tail) {
					var candidates []Point
					if head.x == tail.x || head.y == tail.y {
						candidates = []Point{
							{tail.x + 1, tail.y},
							{tail.x - 1, tail.y},
							{tail.x, tail.y + 1},
							{tail.x, tail.y - 1},
						}
					} else {
						candidates = []Point{
							{tail.x + 1, tail.y + 1},
							{tail.x + 1, tail.y - 1},
							{tail.x - 1, tail.y + 1},
							{tail.x - 1, tail.y - 1},
						}
					}

					dist := head.getManhattanDistance(candidates[0])
					closest := candidates[0]
					for k := 0; k < len(candidates); k++ {
						if head.getManhattanDistance(candidates[k]) <= dist {
							dist = head.getManhattanDistance(candidates[k])
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
	knots []Point
}

func newRope(length int) Rope {
	r := Rope{}
	r.knots = make([]Point, length)

	for i := 0; i < length; i++ {
		r.knots[i] = Point{0, 0}
	}

	return r
}

type Instruction struct {
	direction string
	distance  int
}

type Point struct {
	x int
	y int
}

func addIfNew(ps []Point, newPoint Point) []Point {
	seen := false
	for _, point := range ps {
		seen = seen || point.equals(newPoint)
	}

	if !seen {
		ps = append(ps, newPoint)
	}

	return ps
}

func (p1 Point) equals(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func (p1 Point) tooFarAway(p2 Point) bool {
	xDist := int(math.Abs(float64(p1.x - p2.x)))
	yDist := int(math.Abs(float64(p1.y - p2.y)))

	return xDist > 1 || yDist > 1
}

func (p1 Point) getManhattanDistance(p2 Point) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x))) + int(math.Abs(float64(p1.y)-float64(p2.y)))
}

func (p Point) translate(direction string) Point {
	newP := Point{}
	switch direction {
	case "U":
		newP = Point{p.x, p.y - 1}
	case "D":
		newP = Point{p.x, p.y + 1}
	case "L":
		newP = Point{p.x - 1, p.y}
	case "R":
		newP = Point{p.x + 1, p.y}
	}
	return newP
}
