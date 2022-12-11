package day7

import (
	"strconv"
	"strings"

	"github.com/michael-ryan/AoC22/common"
)

type dir struct {
	files map[string]int
	dirs  map[string]*dir
	size  int
}

func Answers(inputPath string) (int, int) {
	lines := common.ReadFileAsStringArray(inputPath)
	return puzzle1(lines), puzzle2(lines)
}

func puzzle1(lines []string) int {
	if lines[0] != "$ cd /" {
		panic("Not implemented!")
	}

	root := buildTree(lines[1:])

	dirsSmallEnough := root.findDirsSmallerThanN(100000)

	totalSize := 0
	for _, d := range dirsSmallEnough {
		totalSize += d.size
	}

	return totalSize
}

func puzzle2(lines []string) int {
	if lines[0] != "$ cd /" {
		panic("Not implemented!")
	}

	root := buildTree(lines[1:])

	requiredSpace := root.size - 40000000

	dirsBigEnough := root.findDirsBiggerThanN(requiredSpace)

	smallestDir := dirsBigEnough[0]
	for _, d := range dirsBigEnough {
		if d.size < smallestDir.size {
			smallestDir = d
		}
	}

	return smallestDir.size
}

func (root *dir) findDirsSmallerThanN(n int) []*dir {
	var path *common.Stack[*dir] = common.NewStack[*dir]()
	path.Push(root)

	dirs := make([]*dir, 0)

	for _, d := range root.dirs {
		dirs = append(dirs, d.findDirsSmallerThanN(n)...)
	}

	if root.size <= n {
		dirs = append(dirs, root)
	}

	return dirs
}

func (root *dir) findDirsBiggerThanN(n int) []*dir {
	var path *common.Stack[*dir] = common.NewStack[*dir]()
	path.Push(root)

	dirs := make([]*dir, 0)

	for _, d := range root.dirs {
		dirs = append(dirs, d.findDirsBiggerThanN(n)...)
	}

	if root.size >= n {
		dirs = append(dirs, root)
	}

	return dirs
}

func buildTree(lines []string) *dir {
	root := newDir()

	var path *common.Stack[*dir] = common.NewStack[*dir]()
	path.Push(root)

	for _, instruction := range lines {
		splitInstruction := strings.Split(instruction, " ")
		if instruction[0] == '$' {
			// input
			switch splitInstruction[1] {
			case "cd":
				if splitInstruction[2] == ".." {
					path.Pop()
				} else {
					dirname := splitInstruction[2]
					if _, ok := path.Peek().dirs[dirname]; !ok {
						path.Peek().dirs[dirname] = newDir()
					}
					path.Push(path.Peek().dirs[dirname])
				}
			case "ls":
			}
		} else {
			// output
			if splitInstruction[0] == "dir" {
				path.Peek().dirs[splitInstruction[1]] = newDir()
			} else {
				path.Peek().files[splitInstruction[1]], _ = strconv.Atoi(splitInstruction[0])
			}
		}
	}

	root.computeSize()

	return root
}

func newDir() *dir {
	return &dir{
		files: make(map[string]int),
		dirs:  make(map[string]*dir),
		size:  -1,
	}
}

func (d *dir) computeSize() {
	totalSize := 0
	for _, size := range d.files {
		totalSize += size
	}

	for _, subDir := range d.dirs {
		subDir.computeSize()
		totalSize += subDir.size
	}

	d.size = totalSize
}
