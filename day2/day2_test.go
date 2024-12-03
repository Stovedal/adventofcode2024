package day2

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFormatInput(t *testing.T) {
	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1 4 5 2\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n"

	result := FormatInput(input)

	expected := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1, 4, 5, 2}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	if !(utils.MatrixEquals(result, expected)) {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestDecreasingValuesAreConsideredSafe(t *testing.T) {
	input := [][]int{{7, 6, 4, 2, 1}}

	expected := 1

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestIncreasingValuesAreConsideredSafe(t *testing.T) {
	input := [][]int{{1, 3, 5, 6, 8}}

	expected := 1

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestAlternatingDecreasingAndIncreasingValuesAreConsideredUnsafe(t *testing.T) {
	input := [][]int{{1, 3, 2, 6, 9}}

	expected := 0

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestTooSlowlyDecreasingValuesAreConsideredUnsafe(t *testing.T) {
	input := [][]int{{7, 8, 8, 6, 5}}

	expected := 0

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestTooSlowlyIncreasingValuesAreConsideredUnsafe(t *testing.T) {
	input := [][]int{{1, 3, 3, 6, 9}}

	expected := 0

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestTooRapidlyDecreasingValuesAreConsideredUnsafe(t *testing.T) {
	input := [][]int{{11, 8, 3, 2, 1}}

	expected := 0

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestTooRapidlyIncreasingValuesAreConsideredUnsage(t *testing.T) {
	input := [][]int{{1, 3, 4, 10, 11}}

	expected := 0

	result := CountSafeReports(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func TestPanicsOnTooShortReport(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	input := [][]int{{1}}

	CountSafeReports(input)
}

func TestTask1Integration(t *testing.T) {
	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

	expected := 2

	result := Task1(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}
