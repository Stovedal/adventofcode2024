package utils

func ArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func MatrixEquals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !ArrayEquals(v, b[i]) {
			return false
		}
	}

	return true
}

func StringArrayEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func StringMatrixEquals(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !StringArrayEquals(v, b[i]) {
			return false
		}
	}

	return true
}

func ArrayContains(a []int, value int) bool {
	for _, element := range a {
		if element == value {
			return true
		}
	}

	return false
}

func ArrayRemoveElementAt(a []int, index int) []int {
	result := make([]int, len(a)-1)

	copy(result, a[:index])
	copy(result[index:], a[index+1:])

	return result
}
