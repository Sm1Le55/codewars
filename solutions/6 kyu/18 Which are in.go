package kata

import (
	"sort"
	"strings"
)

//InArray ...
func InArray(array1 []string, array2 []string) []string {
	result := make([]string, 0)
	buf := strings.Join(array2, "_")

	for _, element := range array1 {
		if strings.Count(buf, element) > 0 {
			result = append(result, element)
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	for i := 0; i < len(result)-1; i++ {
		if result[i] == result[i+1] {
			result = append(result[:i], result[i+1:]...)
		}
	}

	return result
}
