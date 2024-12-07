package day4

import (
	"regexp"
	"strings"
)

func Task1(input string) int {
	return CountWordsInMatrix("XMAS", FormatInput(input))
}

func Task2(input string) int {
	return CountXmasesInMatrix(FormatInput(input))
}

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

func CountWordsInMatrix(word string, matrix [][]string) int {
	count := 0

	count += countHorizontalWordsInMatrix(word, matrix)

	count += countVerticalWordsInMatrix(word, matrix)

	count += countDiagonalWordsInMatrix(word, matrix)

	return count
}

func countHorizontalWordsInMatrix(word string, matrix [][]string) int {
	count := 0

	for _, horizontalLine := range matrix {
		count += countWordsInLine(word, horizontalLine)
	}

	return count
}

func countVerticalWordsInMatrix(word string, matrix [][]string) int {
	count := 0

	for i := 0; i < len(matrix[0]); i++ {
		verticalLine := []string{}

		for j := 0; j < len(matrix); j++ {
			verticalLine = append(verticalLine, matrix[j][i])
		}

		count += countWordsInLine(word, verticalLine)
	}

	return count
}

func countDiagonalWordsInMatrix(word string, matrix [][]string) int {
	count := 0

	yMax := len(matrix)
	xMax := len(matrix[0])

	for i := -xMax; i < xMax; i++ {
		diagonalLineLtr := []string{}
		diagonalLineRtl := []string{}

		for y := 0; y < yMax; y++ {
			rtlX := y + i

			if rtlX >= 0 && rtlX < xMax {
				value := matrix[y][rtlX]
				diagonalLineLtr = append(diagonalLineLtr, value)
			}

			ltrX := (xMax - 1) - rtlX

			if ltrX >= 0 && ltrX < xMax {
				value := matrix[y][ltrX]
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
