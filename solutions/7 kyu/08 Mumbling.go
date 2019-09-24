package kata

import "strings"

func Accum(s string) string {
    result := ""
    for i,c := range s {
      result += strings.ToUpper(string(c))
      for j := 0; j < i; j++ {
        result += strings.ToLower(string(c))
      }
      result += "-"
    }    
    return strings.TrimRight(result,"-")
}