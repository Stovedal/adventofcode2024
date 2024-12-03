package day2

import (
	"adventofcode2024/utils"
	"math"
	"strconv"
	"strings"
)

func FormatInput(input string) [][]int {
	rows := strings.Split(input, "\n")

	output := [][]int{}

	for _, row := range rows {
		if row == "" {
			continue
		}

		output = append(output, []int{})

		columns := strings.Split(row, " ")

		for _, column := range columns {

			value, err := strconv.Atoi(column)

			if err != nil {
				panic("Could not convert column to int")
			}

			output[len(output)-1] = append(output[len(output)-1], value)

		}
	}
	println("%v", output)

	return output
}

func Task1(input string) int {
	return CountSafeReports(FormatInput(input))
}

func Task2(input string) int {
	return 0
}

func CountSafeReports(reports [][]int) int {
	safeReportCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportCount++
		}
	}

	return safeReportCount
}

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		panic("Report is too short")
	}

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]

		isGraduallyAscendingOrDescending := utils.ArrayContains([]int{1, 2, 3}, int(math.Abs(float64(diff))))
		isContinouslyAscendingOrDescending := report[0]-report[1] > 0 && diff > 0 || report[0]-report[1] < 0 && diff < 0

		if !isGraduallyAscendingOrDescending || !isContinouslyAscendingOrDescending {
			return false
		}
	}

	return true
}
