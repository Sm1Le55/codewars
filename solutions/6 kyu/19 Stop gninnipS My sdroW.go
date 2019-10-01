package kata

import "strings"

//SpinWords ...
func SpinWords(str string) (result string) {
	words := strings.Split(str, " ")
	for _, word := range words {
		if len(word) >= 5 {
			buf := ""
			for i := len(word) - 1; i >= 0; i-- {
				buf += string(word[i])
			}
			result += buf + " "
		} else {
			result += word + " "
		}
	}
	result = strings.TrimSpace(result)
	return result
} // SpinWords
