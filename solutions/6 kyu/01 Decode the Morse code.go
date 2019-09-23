package kata

import "strings"

const WORDS_SEPARATOR = "   "
const CHARS_SEPARATOR = " "

func DecodeMorse(morseCode string) (result string) {
  words := strings.Split(morseCode,WORDS_SEPARATOR)
  
  for _, word := range words {
    chars := strings.Split(word,CHARS_SEPARATOR)
    for _, char := range chars {
      result += MORSE_CODE[char]
    }
    result += " "
  }
  result = strings.Trim(result," ")
  return
}