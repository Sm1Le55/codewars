package kata

import (
	"strings"
	"unicode/utf8"
)

//ValidParentheses ...
func ValidParentheses(parens string) bool {
	parens = findParentheses(parens)
	if len(parens) > 0 {
		return false
	} else {
		return true
	}
}

func findParentheses(parens string) string {
	parens = strings.ReplaceAll(parens, "()", "")
	if utf8.RuneCountInString(parens) < 2 || strings.Count(parens, "()") < 1 {
		return parens
	}
	return findParentheses(parens)
}
