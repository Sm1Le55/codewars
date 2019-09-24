package kata

import "strings"

func solution(str, ending string) bool {
  return strings.Contains(str, ending)  
}