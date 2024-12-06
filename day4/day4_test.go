package day4

import (
	"adventofcode2024/utils"
	"testing"
)

func Test_format_input(t *testing.T) {
	input := "ABCD\nEFGH\nIJKL\n"

	expected := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
	}

	result := FormatInput(input)

	if !utils.StringMatrixEquals(result, expected) {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_count_words_finds_horizontal_words(t *testing.T) {
	input := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "H", "G", "F"},
	}

	expected := 2

	result := CountWordsInMatrix("FGH", input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_count_words_finds_vertical_words(t *testing.T) {
	input := [][]string{
		{"A", "B", "F", "H"},
		{"E", "E", "G", "G"},
		{"I", "H", "H", "F"},
	}

	expected := 2

	result := CountWordsInMatrix("FGH", input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_count_words_finds_left_to_right_diagonal_words(t *testing.T) {
	input := [][]string{
		{"F", "H", "C", "D"},
		{"E", "G", "G", "H"},
		{"I", "H", "H", "F"},
	}

	expected := 2

	result := CountWordsInMatrix("FGH", input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_count_words_finds_right_to_left_diagonal_words(t *testing.T) {
	input := [][]string{
		{"D", "C", "F", "H"},
		{"E", "G", "G", "H"},
		{"H", "F", "G", "F"},
		{"H", "F", "L", "F"},
	}

	expected := 3

	result := CountWordsInMatrix("FGH", input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_count_x_mases_finds_xmases(t *testing.T) {
	input := [][]string{
		{"S", "C", "S", "H"},
		{"E", "A", "G", "H"},
		{"M", "S", "M", "F"},
		{"H", "F", "L", "F"},
	}

	expected := 1

	result := CountXmasesInMatrix(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_integration_test_task_1(t *testing.T) {
	input := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX\n"

	expected := 18

	result := Task1(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_integration_test_task_2(t *testing.T) {
	input := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX\n"

	expected := 9

	result := Task2(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}
