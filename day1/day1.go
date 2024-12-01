package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func FormatInput(input string) ([]int, []int, int) {
	rows := strings.Split(input, "\n")

	leftHandSide := make([]int, len(rows))
	rightHandSide := make([]int, len(rows))

	length := 0

	for i, row := range rows {

		if row == "" {
			continue
		}

		leftHandNumber, err := strconv.Atoi(strings.Split(row, "   ")[0])

		if err != nil {
			panic(err)
		}

		rightHandNumber, err := strconv.Atoi(strings.Split(row, "   ")[1])

		if err != nil {
			panic(err)
		}

		leftHandSide[i] = leftHandNumber
		rightHandSide[i] = rightHandNumber

		length++
	}

	return leftHandSide, rightHandSide, length
}

func CalculateDifferenceSum(leftHandSide []int, rightHandSide []int, length int) int {
	sort.Slice(leftHandSide, func(i, j int) bool {
		return leftHandSide[i] < leftHandSide[j]
	})

	sort.Slice(rightHandSide, func(i, j int) bool {
		return rightHandSide[i] < rightHandSide[j]
	})

	difference_sum := 0

	for i := 0; i < length; i++ {
		difference_sum += int(math.Abs(float64(rightHandSide[i] - leftHandSide[i])))
	}

	return difference_sum
}

func CalculateSimilarityScore(leftHandSide []int, rightHandSide []int, length int) int {
	similarity_scores := make(map[int]int)

	total_similarity_score := 0

	for i := 0; i < length; i++ {
		similarity_score, exists := similarity_scores[leftHandSide[i]]

		if exists {
			total_similarity_score += similarity_score
		} else {
			similarity_score := 0
			for j := 0; j < length; j++ {
				if leftHandSide[i] == rightHandSide[j] {
					similarity_score += leftHandSide[i]
				}
			}
			similarity_scores[leftHandSide[i]] = similarity_score

			total_similarity_score += similarity_score
		}

	}

	return total_similarity_score
}

func Task1(input string) int {
	return CalculateDifferenceSum(FormatInput(input))
}

func Task2(input string) int {
	return CalculateSimilarityScore(FormatInput(input))
}
