package kata

import "strings"

func Solution(str string) (result []string) {
	if len(str)%2 != 0 {
		str = str + "_"
	}
	s := strings.Split(str, "")

	for i := 0; i < len(s)-1; i += 2 {
		result = append(result, s[i]+s[i+1])
	}

	return result
}
