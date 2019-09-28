package kata

import (
	"strings"
)

// You may get original char by morse code like this: MORSE_CODE[char]

//DecodeBits ...
func DecodeBits(bits string) string {

	bits = strings.Trim(bits, "0")

	time := searchTime(bits)
	k0, k1 := "0", "1"

	for i := 0; i < time; i++ {
		k0 += "0"
		k1 += "1"
	}

	bits = strings.ReplaceAll(bits, k0, "0")
	bits = strings.ReplaceAll(bits, k1, "1")
	bits = strings.ReplaceAll(bits, "0", " ")
	bits = strings.ReplaceAll(bits, "111", "-")
	bits = strings.ReplaceAll(bits, "1", ".")

	return bits
}

//DecodeMorse ...
func DecodeMorse(morseCode string) string {
	result := ""
	words := strings.Split(morseCode, "       ")
	for _, word := range words {
		chars := strings.Split(word, "   ")
		for _, ch := range chars {
			result += MORSE_CODE[strings.ReplaceAll(ch, " ", "")]
		}
		result += " "
	}
	result = strings.TrimSpace(result)
	return result
}

//searchTime ...
func searchTime(bits string) int {

	buf := strings.Split(bits, "")
	min := len(buf) - 1
	length := 0

	for i := 0; i < len(buf)-1; i++ {
		if buf[i] == buf[i+1] {
			length++
		} else {
			if length < min {
				min = length
			}
			length = 0
		}
	}

	return min
}
