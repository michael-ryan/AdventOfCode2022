package main

import (
	"fmt"

	day1 "github.com/michael-ryan/AoC22/Day1"
	day2 "github.com/michael-ryan/AoC22/Day2"
	day3 "github.com/michael-ryan/AoC22/Day3"
	day4 "github.com/michael-ryan/AoC22/Day4"
	day5 "github.com/michael-ryan/AoC22/Day5"
)

func main() {
	showDay1Results()
	showDay2Results()
	showDay3Results()
	showDay4Results()
	showDay5Results()
}

func showDay1Results() {
	fmt.Println("Day 1")
	fmt.Println("Puzzle 1:", day1.Puzzle1())
	fmt.Println("Puzzle 2:", day1.Puzzle2())
}

func showDay2Results() {
	fmt.Println("Day 2")
	fmt.Println("Puzzle 1:", day2.Puzzle1())
	fmt.Println("Puzzle 2:", day2.Puzzle2())
}

func showDay3Results() {
	fmt.Println("Day 3")
	fmt.Println("Puzzle 1:", day3.Puzzle1())
	fmt.Println("Puzzle 2:", day3.Puzzle2())
}

func showDay4Results() {
	fmt.Println("Day 4")
	fmt.Println("Puzzle 1:", day4.Puzzle1())
	fmt.Println("Puzzle 2:", day4.Puzzle2())
}

func showDay5Results() {
	fmt.Println("Day 5")
	fmt.Println("Puzzle 1:", day5.Puzzle1())
	fmt.Println("Puzzle 2:", day5.Puzzle2())
}
