package kata

import "strings"

func Solution(word string) string {    
    var b strings.Builder
    for i:=len(word)-1; i>-1; i-- {
      b.WriteByte(word[i])	    
    }
    return b.String()
}