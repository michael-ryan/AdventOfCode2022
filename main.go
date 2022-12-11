package main

import (
	"fmt"

	day1 "github.com/michael-ryan/AoC22/Day1"
	day2 "github.com/michael-ryan/AoC22/Day2"
	day3 "github.com/michael-ryan/AoC22/Day3"
	day4 "github.com/michael-ryan/AoC22/Day4"
	day5 "github.com/michael-ryan/AoC22/Day5"
	/*
		day6 "github.com/michael-ryan/AoC22/Day6"
		day7 "github.com/michael-ryan/AoC22/Day7"
		day8 "github.com/michael-ryan/AoC22/Day8"
		day9 "github.com/michael-ryan/AoC22/Day9"
		day10 "github.com/michael-ryan/AoC22/Day10"
		day11 "github.com/michael-ryan/AoC22/Day11"
		day12 "github.com/michael-ryan/AoC22/Day12"
		day13 "github.com/michael-ryan/AoC22/Day13"
		day14 "github.com/michael-ryan/AoC22/Day14"
		day15 "github.com/michael-ryan/AoC22/Day15"
		day16 "github.com/michael-ryan/AoC22/Day16"
		day17 "github.com/michael-ryan/AoC22/Day17"
		day18 "github.com/michael-ryan/AoC22/Day18"
		day19 "github.com/michael-ryan/AoC22/Day19"
		day20 "github.com/michael-ryan/AoC22/Day20"
		day21 "github.com/michael-ryan/AoC22/Day21"
		day22 "github.com/michael-ryan/AoC22/Day22"
		day23 "github.com/michael-ryan/AoC22/Day23"
		day24 "github.com/michael-ryan/AoC22/Day24"
		day25 "github.com/michael-ryan/AoC22/Day25"
	*/)

func main() {
	showDay1Results()
	showDay2Results()
	showDay3Results()
	showDay4Results()
	showDay5Results()
       /*
	        showDay6Results()
		showDay7Results()
		showDay8Results()
		showDay9Results()
		showDay10Results()
		showDay11Results()
		showDay12Results()
		showDay13Results()
		showDay14Results()
		showDay15Results()
		showDay16Results()
		showDay17Results()
		showDay18Results()
		showDay19Results()
		showDay20Results()
		showDay21Results()
		showDay22Results()
		showDay23Results()
		showDay24Results()
		showDay25Results()
	*/
}

func showDay1Results() {
	fmt.Println("Day 1")
	one, two := day1.Answers("resources/inputs/day1.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}

func showDay2Results() {
	fmt.Println("Day 2")
	one, two := day2.Answers("resources/inputs/day2.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}

func showDay3Results() {
	fmt.Println("Day 3")
	one, two := day3.Answers("resources/inputs/day3.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}

func showDay4Results() {
	fmt.Println("Day 4")
	one, two := day4.Answers("resources/inputs/day4.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}

func showDay5Results() {
	fmt.Println("Day 5")
	one, two := day5.Answers("resources/inputs/day5.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}

/*
func showDay6Results() {
	fmt.Println("Day 6")
	one, two := day6.Answers("resources/inputs/day6.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay7Results() {
	fmt.Println("Day 7")
	one, two := day7.Answers("resources/inputs/day7.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay8Results() {
	fmt.Println("Day 8")
	one, two := day8.Answers("resources/inputs/day8.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay9Results() {
	fmt.Println("Day 9")
	one, two := day9.Answers("resources/inputs/day9.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay10Results() {
	fmt.Println("Day 10")
	one, two := day10.Answers("resources/inputs/day10.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay11Results() {
	fmt.Println("Day 11")
	one, two := day11.Answers("resources/inputs/day11.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay12Results() {
	fmt.Println("Day 12")
	one, two := day12.Answers("resources/inputs/day12.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay13Results() {
	fmt.Println("Day 13")
	one, two := day13.Answers("resources/inputs/day13.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay14Results() {
	fmt.Println("Day 14")
	one, two := day14.Answers("resources/inputs/day14.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay15Results() {
	fmt.Println("Day 15")
	one, two := day15.Answers("resources/inputs/day15.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay16Results() {
	fmt.Println("Day 16")
	one, two := day16.Answers("resources/inputs/day16.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay17Results() {
	fmt.Println("Day 17")
	one, two := day17.Answers("resources/inputs/day17.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay18Results() {
	fmt.Println("Day 18")
	one, two := day18.Answers("resources/inputs/day18.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay19Results() {
	fmt.Println("Day 19")
	one, two := day19.Answers("resources/inputs/day19.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay20Results() {
	fmt.Println("Day 20")
	one, two := day20.Answers("resources/inputs/day20.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay21Results() {
	fmt.Println("Day 21")
	one, two := day21.Answers("resources/inputs/day21.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay22Results() {
	fmt.Println("Day 22")
	one, two := day22.Answers("resources/inputs/day22.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay23Results() {
	fmt.Println("Day 23")
	one, two := day23.Answers("resources/inputs/day23.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay24Results() {
	fmt.Println("Day 24")
	one, two := day24.Answers("resources/inputs/day24.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
func showDay25Results() {
	fmt.Println("Day 25")
	one, two := day25.Answers("resources/inputs/day25.txt")
	fmt.Println("Puzzle 1:", one)
	fmt.Println("Puzzle 2:", two)
}
*/
