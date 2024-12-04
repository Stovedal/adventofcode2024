package day2

import (
	"adventofcode2024/utils"
	"testing"
)

func Test_format_input(t *testing.T) {
	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1 4 5 2\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n"

	result := FormatInput(input)

	expected := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1, 4, 5, 2}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	if !(utils.MatrixEquals(result, expected)) {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_decreasing_values_are_considered_safe(t *testing.T) {
	input := [][]int{{7, 6, 4, 2, 1}}

	expected := 1

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_increasing_values_are_considered_safe(t *testing.T) {
	input := [][]int{{1, 3, 5, 6, 8}}

	expected := 1

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_alternating_decreasing_and_increasing_values_are_considered_unsafe(t *testing.T) {
	input := [][]int{{1, 3, 2, 6, 9}}

	expected := 0

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_slowly_decreasing_values_are_considered_unsafe(t *testing.T) {
	input := [][]int{{7, 8, 8, 6, 5}}

	expected := 0

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_slowly_increasing_value_sare_considered_unsafe(t *testing.T) {
	input := [][]int{{1, 3, 3, 6, 9}}

	expected := 0

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_rapidly_decreasing_values_are_considered_unsafe(t *testing.T) {
	input := [][]int{{11, 8, 3, 2, 1}}

	expected := 0

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_rapidly_increasing_values_are_considered_unsafe(t *testing.T) {
	input := [][]int{{1, 3, 4, 10, 11}}

	expected := 0

	result := CountSafeReports(input, false)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_rapidly_increasing_values__once_are_considered_safe_when_dampened(t *testing.T) {
	input := [][]int{{1, 3, 4, 6, 11}}

	expected := 1

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_alternating_decreasing_and_increasing_values_once_are_considered_safe_when_dampened(t *testing.T) {
	input := [][]int{{1, 3, 2, 6, 9}}

	expected := 1

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_slowly_decreasing_values_once_are_considered_safe_when_dampened(t *testing.T) {
	input := [][]int{{7, 6, 5, 5, 4}}

	expected := 1

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_slowly_increasing_values_once_are_considered_safe_when_dampened(t *testing.T) {
	input := [][]int{{1, 3, 3, 6, 9}}

	expected := 1

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_rapidly_decreasing_values_once_are_considered_safe_when_dampened(t *testing.T) {
	input := [][]int{{11, 4, 3, 2, 1}}

	expected := 1

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_rapidly_increasing_values_twice_are_considered_unsafe_when_dampened(t *testing.T) {
	input := [][]int{{1, 3, 4, 10, 16}}

	expected := 0

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_alternating_decreasing_and_increasing_values_twice_are_considered_unsafe_when_dampened(t *testing.T) {
	input := [][]int{{1, 3, 2, 6, 5, 6}}

	expected := 0

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_slowly_decreasing_values_twice_are_considered_unsafe_when_dampened(t *testing.T) {
	input := [][]int{{7, 8, 8, 6, 6}}

	expected := 0

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_slowly_increasing_values_twice_are_considered_unsafe_when_dampened(t *testing.T) {
	input := [][]int{{1, 3, 3, 6, 6}}

	expected := 0

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_too_rapidly_decreasing_values_twice_are_considered_unsafe_when_dampened(t *testing.T) {
	input := [][]int{{18, 8, 3, 2, 1}}

	expected := 0

	result := CountSafeReports(input, true)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}

func Test_panics_on_too_short_report(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	input := [][]int{{1}}

	CountSafeReports(input, false)
}

func Test_task1_integration(t *testing.T) {
	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

	expected := 2

	result := Task1(input)

	if result != expected {
		t.Errorf("Test input %v, resulted in %v instead of %v", input, result, expected)
	}
}
