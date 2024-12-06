package day3

import (
	"regexp"
	"strconv"
)

type Instruction struct {
	operation string
	arguments []int
}

func FormatInput(input string) []Instruction {

	instructions := []Instruction{}

	instructionRegex, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	operationRegex, _ := regexp.Compile(`mul|don't|do`)
	digitRegex, _ := regexp.Compile(`\d+`)

	matches := instructionRegex.FindAllString(input, -1)

	for _, match := range matches {

		operation := operationRegex.FindString(match)

		argumentsStrings := digitRegex.FindAllString(match, -1)

		arguments := []int{}

		for _, argument := range argumentsStrings {
			argumentInt, err := strconv.Atoi(argument)

			if err != nil {
				panic(err)
			}

			arguments = append(arguments, argumentInt)
		}

		instructions = append(instructions, Instruction{operation: operation, arguments: arguments})
	}

	return instructions
}

func ExecuteInstructions(instructions []Instruction, withEnabledState bool) int {
	result := 0
	enabled := true

	for _, instruction := range instructions {
		if instruction.operation == "mul" && enabled {
			result += instruction.arguments[0] * instruction.arguments[1]
		}

		if withEnabledState {
			if instruction.operation == "do" {
				enabled = true
			} else if instruction.operation == "don't" {
				enabled = false
			}
		}
	}

	return result
}

func Task1(input string) int {
	return ExecuteInstructions(FormatInput(input), false)
}

func Task2(input string) int {
	return ExecuteInstructions(FormatInput(input), true)
}
