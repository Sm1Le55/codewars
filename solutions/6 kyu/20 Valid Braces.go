package kata

import "strings"

//ValidBraces ...
func ValidBraces(str string) bool {
	for checkBracesPair(str) {
		str = strings.ReplaceAll(str, "()", "")
		str = strings.ReplaceAll(str, "[]", "")
		str = strings.ReplaceAll(str, "{}", "")
	}
	if len(str) > 0 {
		return false
	} else {
		return true
	}
}

//checkBracesPair ...
func checkBracesPair(str string) bool {
	braces1 := strings.Count(str, "()")
	braces2 := strings.Count(str, "[]")
	braces3 := strings.Count(str, "{}")

	return braces1 > 0 || braces2 > 0 || braces3 > 0
}
