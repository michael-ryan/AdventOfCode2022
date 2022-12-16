package common

import (
	"math"
	"os"
	"strings"
)

func ReadFileAsStringArray(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func Stuff(data string) []string {
	return strings.Split(string(data), "\n")
}

type Point struct {
	X int
	Y int
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func (p1 Point) GetManhattanDistance(p2 Point) int {
	return int(math.Abs(float64(p1.X)-float64(p2.X))) + int(math.Abs(float64(p1.Y)-float64(p2.Y)))
}

func (p Point) Translate(direction string) Point {
	newP := Point{}
	switch direction {
	case "U":
		newP = Point{X: p.X, Y: p.Y - 1}
	case "D":
		newP = Point{X: p.X, Y: p.Y + 1}
	case "L":
		newP = Point{X: p.X - 1, Y: p.Y}
	case "R":
		newP = Point{X: p.X + 1, Y: p.Y}
	}
	return newP
}
