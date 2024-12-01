package day1

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFormatInput(t *testing.T) {
	input := "3  4\n4  3\n2  5\n1  3\n3  9\n3  3"

	leftHandSide, rightHandSide, length := FormatInput(input)

	if !utils.ArrayEquals(leftHandSide, []int{3, 4, 2, 1, 3, 3}) {
		t.Errorf("Test input %q, resulted in %v instead of %v", input, leftHandSide, []int{3, 4, 2, 1, 3, 3})
	}

	if !utils.ArrayEquals(rightHandSide, []int{4, 3, 5, 3, 9, 3}) {
		t.Errorf("Test input %q, resulted in %v instead of %v", input, rightHandSide, []int{4, 3, 5, 3, 9, 3})
	}

	if length != 6 {
		t.Errorf("Test input %q, resulted in %v instead of %v", input, length, 6)
	}
}

func TestCalculatesDifferenceBetweenLists(t *testing.T) {
	leftHandSide := []int{3, 4, 2, 1, 3, 3}
	rightHandSide := []int{4, 3, 5, 3, 9, 3}
	length := 6

	result := CalculateDifferenceSum(leftHandSide, rightHandSide, length)
	expected := 11

	if result != expected {
		t.Errorf("Calculating difference between lists resulted in %v instead of %v", result, expected)
	}
}

func TestCalculatesDifferenceBetweenListsOtherWayAround(t *testing.T) {
	leftHandSide := []int{4, 3, 5, 3, 9, 3}
	rightHandSide := []int{3, 4, 2, 1, 3, 3}
	length := 6

	result := CalculateDifferenceSum(leftHandSide, rightHandSide, length)
	expected := 11

	if result != expected {
		t.Errorf("Calculating difference between lists other way around resulted in %v instead of %v", result, expected)
	}
}

func TestCalculateSimilarityScoreBetweenLists(t *testing.T) {
	leftHandSide := []int{3, 4, 2, 1, 3, 3}
	rightHandSide := []int{4, 3, 5, 3, 9, 3}
	length := 6

	result := CalculateSimilarityScore(leftHandSide, rightHandSide, length)
	expected := 31

	if result != expected {
		t.Errorf("Calculating similarity score between lists resulted in %v instead of %v", result, expected)
	}
}

func TestCalculateSimilarityScoreBetweenListsOtherWayAround(t *testing.T) {

	leftHandSide := []int{4, 3, 5, 3, 9, 3}
	rightHandSide := []int{3, 4, 2, 1, 3, 3}
	length := 6

	result := CalculateSimilarityScore(leftHandSide, rightHandSide, length)
	expected := 31

	if result != expected {
		t.Errorf("Calculating similarity score between lists other way around resulted in %v instead of %v", result, expected)
	}
}
