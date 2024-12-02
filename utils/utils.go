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

func ArrayContains(a []int, value int) bool {
	for _, element := range a {
		if element == value {
			return true
		}
	}

	return false
}
