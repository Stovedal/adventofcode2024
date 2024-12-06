package day3

import (
	"adventofcode2024/utils"
	"testing"
)

func Test_format_input_ignores_corrupt_data(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	expected := []Instruction{
		{operation: "mul", arguments: []int{2, 4}},
		{operation: "mul", arguments: []int{5, 5}},
		{operation: "mul", arguments: []int{11, 8}},
		{operation: "mul", arguments: []int{8, 5}},
	}

	result := FormatInput(input)

	if len(result) != len(expected) {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}

	for i, instruction := range result {
		if instruction.operation != expected[i].operation || !utils.ArrayEquals(instruction.arguments, expected[i].arguments) {
			t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
		}
	}
}

func Test_format_input_registers_operations(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	expected := []Instruction{
		{operation: "mul", arguments: []int{2, 4}},
		{operation: "don't", arguments: []int{}},
		{operation: "mul", arguments: []int{5, 5}},
		{operation: "mul", arguments: []int{11, 8}},
		{operation: "do", arguments: []int{}},
		{operation: "mul", arguments: []int{8, 5}},
	}

	result := FormatInput(input)

	if len(result) != len(expected) {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}

	for i, instruction := range result {
		if instruction.operation != expected[i].operation || !utils.ArrayEquals(instruction.arguments, expected[i].arguments) {
			t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
		}
	}
}

func Test_executing_instructions_multiplies_arguments_for_mul_operation(t *testing.T) {
	input := []Instruction{
		{operation: "mul", arguments: []int{2, 4}},
		{operation: "mul", arguments: []int{5, 5}},
		{operation: "mul", arguments: []int{11, 8}},
		{operation: "mul", arguments: []int{8, 5}},
	}

	expected := 2*4 + 5*5 + 11*8 + 8*5

	result := ExecuteInstructions(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_integration_test_task_1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	expected := 2*4 + 5*5 + 11*8 + 8*5

	result := Task1(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}
