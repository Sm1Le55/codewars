package kata

import (
	"sort"
)

//ValidateSolution ...
func ValidateSolution(m [][]int) bool {
	return ValidateRows(m) && ValidateColumns(m) && ValidateSections(m)
}

//ValidateRows ...
func ValidateRows(m [][]int) bool {
	for i := 0; i < len(m); i++ {
		if len(m[i]) != 9 {
			return false
		}

		buf := make([]int, len(m[i]), len(m[i]))
		copy(buf, m[i])

		if !CheckSlice(buf) {
			return false
		}
	}
	return true
}

//ValidateColumns ...
func ValidateColumns(m [][]int) bool {

	if len(m) != 9 {
		return false
	}

	for i := 0; i < len(m); i++ {

		buf := make([]int, 0, len(m[i]))

		for j := 0; j < len(m); j++ {
			buf = append(buf, m[j][i])
		}

		if !CheckSlice(buf) {
			return false
		}
	}
	return true
}

//ValidateSections ...
func ValidateSections(m [][]int) bool {

	for startIdx := 0; startIdx < len(m); startIdx += 3 {
		for startIdxJ := 0; startIdxJ < len(m); startIdxJ += 3 {

			buf := make([]int, 0, len(m))
			for i := startIdx; i < startIdx+3; i++ {
				for j := startIdxJ; j < startIdxJ+3; j++ {
					buf = append(buf, m[i][j])
				}
			}

			if !CheckSlice(buf) {
				return false
			}
		}
	}
	return true
}

//CheckSlice ...
func CheckSlice(m []int) bool {
	sort.Slice(m, func(i, j int) bool { return m[i] < m[j] })

	for idx := 0; idx < len(m); idx++ {
		if m[idx] == idx+1 {
			continue
		}
		return false
	}
	return true
}
