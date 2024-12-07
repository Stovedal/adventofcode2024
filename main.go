package main

import (
	"adventofcode2024/day1"
	"adventofcode2024/day2"
	"adventofcode2024/day3"
	"adventofcode2024/day4"
	"adventofcode2024/day5"
	"adventofcode2024/input"
	"adventofcode2024/utils"
	"fmt"
)

type Solution struct {
	Task1 func(string) int
	Task2 func(string) int
}

func main() {
	solutions := map[int]Solution{
		1: {Task1: day1.Task1, Task2: day1.Task2},
		2: {Task1: day2.Task1, Task2: day2.Task2},
		3: {Task1: day3.Task1, Task2: day3.Task2},
		4: {Task1: day4.Task1, Task2: day4.Task2},
		5: {Task1: day5.Task1, Task2: day5.Task2},
	}

	solutionIdentifier, sessionToken := utils.ParseUserArguments(len(solutions))

	fmt.Printf("Solving day %v\n", solutionIdentifier)

	input := input.FetchInputForSolution(solutionIdentifier, sessionToken)

	fmt.Printf("\tAnswer to task 1: %v\n", solutions[solutionIdentifier].Task1(input))
	fmt.Printf("\tAnswer to task 2: %v\n", solutions[solutionIdentifier].Task2(input))
}
