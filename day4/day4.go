package day4

import (
	"regexp"
	"strings"
)

func FormatInput(input string) [][]string {
	output := [][]string{}

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}

		output = append(output, strings.Split(row, ""))
	}

	return output
}

func CountXmasesInMatrix(matrix [][]string) int {
	count := 0

	for y := 1; y < len(matrix[0])-1; y++ {
		for x := 1; x < len(matrix)-1; x++ {
			if matrix[x][y] == "A" {
				corners := []string{
					matrix[x-1][y-1],
					matrix[x+1][y+1],
					matrix[x+1][y-1],
					matrix[x-1][y+1],
				}

				sCount := 0
				mCount := 0

				for _, corner := range corners {
					if corner == "S" {
						sCount++
					} else if corner == "M" {
						mCount++
					}
				}

				if sCount == 2 && mCount == 2 {
					count++
				}

			}
		}
	}

	return count
}

func CountWordsInMatrix(word string, matrix [][]string) int {

	count := 0

	for _, horizontalLine := range matrix {
		count += countWordsInLine(word, horizontalLine)
	}

	for i := 0; i < len(matrix[0]); i++ {
		verticalLine := []string{}

		for j := 0; j < len(matrix); j++ {
			verticalLine = append(verticalLine, matrix[j][i])
		}

		count += countWordsInLine(word, verticalLine)
	}

	for k := -len(matrix[0]); k < len(matrix[0]); k++ {
		diagonalLineLtr := []string{}
		diagonalLineRtl := []string{}

		for y := 0; y < len(matrix); y++ {
			if y+k < len(matrix[0]) && y+k >= 0 {
				value := matrix[y][y+k]
				diagonalLineLtr = append(diagonalLineLtr, value)
			}

			if (len(matrix[0])-1)-y-k >= 0 && (len(matrix[0])-1)-y-k < len(matrix[0]) {
				value := matrix[y][(len(matrix[0])-1)-y-k]
				diagonalLineRtl = append(diagonalLineRtl, value)
			}

		}

		count += countWordsInLine(word, diagonalLineLtr)
		count += countWordsInLine(word, diagonalLineRtl)
	}

	return count
}

func countWordsInLine(word string, line []string) int {
	count := 0

	wordRegex := regexp.MustCompile(word)
	reversedWordRegex := regexp.MustCompile(reverseString(word))

	count += len(wordRegex.FindAllString(strings.Join(line, ""), -1))
	count += len(reversedWordRegex.FindAllString(strings.Join(line, ""), -1))

	return count
}

func reverseString(s string) string {
	output := []string{}

	for _, row := range strings.Split(s, "") {
		output = append([]string{row}, output...)
	}

	return strings.Join(output, "")
}

func Task1(input string) int {
	return CountWordsInMatrix("XMAS", FormatInput(input))
}

func Task2(input string) int {
	return CountXmasesInMatrix(FormatInput(input))
}
