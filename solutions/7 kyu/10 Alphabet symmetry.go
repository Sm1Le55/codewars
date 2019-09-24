package kata

import "strings"

func solve(slice []string) []int {
  var result []int
  
  for _, val := range slice {
    cnt := 0
    for idx, char := range strings.ToLower(val) {      
      if (idx==int(char-'a')) {
        cnt++
      }
    }
    result = append(result, cnt)
  }
  return result
}