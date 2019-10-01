package kata

import "strings"

//toWeirdCase ...
func toWeirdCase(str string) string {
	str = strings.ToLower(str)
	words := strings.Split(str, " ")

	for idx, word := range words {
		for i := 0; i < len(word); i++ {
			if i == 0 || i%2 == 0 {
				word = word[:i] + strings.ToUpper(string(word[i])) + word[i+1:]
			}
		}
		words[idx] = word
	}

	result := strings.Join(words, " ")
	return result
}
