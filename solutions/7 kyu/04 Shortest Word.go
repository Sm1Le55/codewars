package kata

import "strings"

func FindShort(s string) int {

	words := strings.Split(s, " ")
	minWordLength := len(s)

	for _, word := range words {
		if len(word) < minWordLength {
			minWordLength = len(word)
		}
	}
	return minWordLength
}
