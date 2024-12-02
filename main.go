package main

import (
	"adventofcode2024/day1"
	"adventofcode2024/day2"
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
	}

	solutionIdentifier, sessionToken := utils.ParseUserArguments(len(solutions))

	fmt.Printf("Solving day %v\n", solutionIdentifier)

	input := input.FetchInputForSolution(solutionIdentifier, sessionToken)

	fmt.Printf("\tAnswer to task 1: %v\n", solutions[solutionIdentifier].Task1(input))
	fmt.Printf("\tAnswer to task 2: %v\n", solutions[solutionIdentifier].Task2(input))
}
