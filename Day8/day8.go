package day8

import (
	"strconv"

	"github.com/michael-ryan/AoC22/common"
)

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func puzzle1(lines []string) int {
	grid := parseGrid(lines)

	resolveVisibilities(grid)

	numberOfVisibleTrees := 0
	for _, row := range grid {
		for _, tree := range row {
			if tree.visible {
				numberOfVisibleTrees++
			}
		}
	}

	return numberOfVisibleTrees
}

func puzzle2(lines []string) int {
	grid := parseGrid(lines)

	resolveScenicScores(grid)

	maxScore := -1
	for _, row := range grid {
		for _, tree := range row {
			if tree.scenicScore > maxScore {
				maxScore = tree.scenicScore
			}
		}
	}

	return maxScore
}

func parseGrid(lines []string) [][]*Tree {
	// grid[y][x] - (0, 0) is top left
	grid := make([][]*Tree, 0)
	for y, line := range lines {
		grid = append(grid, make([]*Tree, 0))
		for _, h := range line {
			height, _ := strconv.Atoi(string(h))
			grid[y] = append(grid[y], newTree(height))
		}
	}

	return grid
}

func resolveScenicScores(grid [][]*Tree) {
	for y := range grid {
		for x, tree := range grid[y] {
			tree.scenicScore = resolveScenicScore(grid, x, y)
		}
	}
}

func resolveScenicScore(grid [][]*Tree, treeX int, treeY int) int {
	myHeight := grid[treeY][treeX].height

	// UP
	up := 0
	for y := treeY - 1; y >= 0; y-- {
		up++
		if grid[y][treeX].height >= myHeight {
			break
		}
	}

	// DOWN
	down := 0
	for y := treeY + 1; y < len(grid); y++ {
		down++
		if grid[y][treeX].height >= myHeight {
			break
		}
	}

	// LEFT
	left := 0
	for x := treeX - 1; x >= 0; x-- {
		left++
		if grid[treeY][x].height >= myHeight {
			break
		}
	}

	// RIGHT
	right := 0
	for x := treeX + 1; x < len(grid[0]); x++ {
		right++
		if grid[treeY][x].height >= myHeight {
			break
		}
	}

	return up * down * left * right
}

func resolveVisibilities(grid [][]*Tree) {
	// going up
	for x := 0; x < len(grid[0]); x++ {
		resolveVisibility(grid, x, len(grid)-1, "UP")
	}

	// going down
	for x := 0; x < len(grid[0]); x++ {
		resolveVisibility(grid, x, 0, "DOWN")
	}

	// going left
	for y := 0; y < len(grid); y++ {
		resolveVisibility(grid, len(grid[0])-1, y, "LEFT")
	}

	// going right
	for y := 0; y < len(grid); y++ {
		resolveVisibility(grid, 0, y, "RIGHT")
	}
}

func resolveVisibility(grid [][]*Tree, initialX int, initialY int, direction string) {
	x := initialX
	y := initialY

	tallestTree := -1
	for {
		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
			break
		}

		if grid[y][x].height > tallestTree {
			tallestTree = grid[y][x].height
			grid[y][x].visible = true
		}

		switch direction {
		case "UP":
			y -= 1
		case "DOWN":
			y += 1
		case "LEFT":
			x -= 1
		case "RIGHT":
			x += 1
		}
	}
}

type Tree struct {
	height      int
	visible     bool
	scenicScore int
}

func newTree(height int) *Tree {
	return &Tree{
		height:      height,
		visible:     false,
		scenicScore: 0,
	}
}
