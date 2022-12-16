package day12

import (
	"math"

	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func puzzle1(lines []string) int {
	heightMap := convertInputIntoRuneGrid(lines)

	graph := buildGraph(heightMap, false)

	start := findCoordinateOfSpecificRune(heightMap, 'S')
	end := findCoordinateOfSpecificRune(heightMap, 'E')

	distances := graph.Dijkstra(start)

	return distances[end]
}

func puzzle2(lines []string) int {
	heightMap := convertInputIntoRuneGrid(lines)

	end := findCoordinateOfSpecificRune(heightMap, 'E')

	graph := buildGraph(heightMap, true)
	distances := graph.Dijkstra(end)

	candidates := make([]common.Point, 0)

	for point := range distances {
		if heightMap[point.Y][point.X] == 'S' || heightMap[point.Y][point.X] == 'a' {
			candidates = append(candidates, point)
		}
	}

	lowestSteps := math.MaxInt64
	for _, point := range candidates {
		if distances[point] < lowestSteps {
			lowestSteps = distances[point]
		}
	}

	return lowestSteps
}

func findCoordinateOfSpecificRune(heightMap [][]rune, target rune) common.Point {
	var coordinate common.Point

	for y := range heightMap {
		for x := range heightMap[0] {
			if heightMap[y][x] == target {
				coordinate = common.Point{X: x, Y: y}
			}
		}
	}

	return coordinate
}

func findCoordinatesOfSpecificRune(heightMap [][]rune, target rune) []common.Point {
	coordinates := make([]common.Point, 0)

	for y := range heightMap {
		for x := range heightMap[0] {
			if heightMap[y][x] == target {
				coordinates = append(coordinates, common.Point{X: x, Y: y})
			}
		}
	}

	return coordinates
}

func convertInputIntoRuneGrid(lines []string) [][]rune {
	heightMap := make([][]rune, 0)

	for _, line := range lines {
		heightMap = append(heightMap, []rune(line))
	}

	return heightMap
}

func buildGraph(heightMap [][]rune, backwards bool) common.Graph[common.Point] {
	graph := common.Graph[common.Point]{}
	for y := range heightMap {
		for x := range heightMap[y] {
			targets := getValidTargetsForEdges(heightMap, x, y, backwards)
			graph[common.Point{X: x, Y: y}] = createSliceOfEdges(targets)
		}
	}

	return graph
}

func getValidTargetsForEdges(heightMap [][]rune, x int, y int, backwards bool) []common.Point {
	targets := make([]common.Point, 0)
	for _, d := range [...]string{"U", "D", "L", "R"} {
		candidate := common.Point{X: x, Y: y}.Translate(d)

		if candidate.X >= 0 && candidate.X < len(heightMap[0]) && candidate.Y >= 0 && candidate.Y < len(heightMap) {
			from := heightMap[y][x]
			to := heightMap[candidate.Y][candidate.X]
			if from == 'S' {
				from = 'a'
			} else if from == 'E' {
				from = 'z'
			}

			if to == 'S' {
				to = 'a'
			} else if to == 'E' {
				to = 'z'
			}
			if !backwards {
				if to < from || to == from || to == from+1 {
					targets = append(targets, candidate)
				}
			} else {
				if to > from || to == from || to == from-1 {
					targets = append(targets, candidate)
				}
			}

		}
	}

	return targets
}

func createSliceOfEdges(points []common.Point) []common.Edge[common.Point] {
	edges := make([]common.Edge[common.Point], 0)

	for _, point := range points {
		edges = append(edges, common.Edge[common.Point]{To: point, Weight: 1})
	}

	return edges
}
