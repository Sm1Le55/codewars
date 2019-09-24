package kata

import "strings"

func Scale(s string, k, n int ) (result string) {
    parts := strings.Split(s,"\n")
    for _,part := range parts {
      for i:=0; i<n; i++ {
        for _,c := range part {
          for j:=0; j<k; j++ {
            result += string(c)
          }
        }
        result += "\n"
      }      
    }
    result = strings.TrimRight(result,"\n")
    return
}