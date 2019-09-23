package kata

import "strings"

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(str, vowel)
	}

	return count
}
