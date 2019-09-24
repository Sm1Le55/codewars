package kata

import "strings"

func duplicate_count(s1 string) int {
    duplicates := 0    
    
    s1 = strings.ToUpper(s1)    
    chars := strings.Split(s1, "")
    
    for i:=0; i<len(chars)-1; i++ {      
      cnt := 1
      for j:=i+1; j<len(chars); j++ {
        if chars[j]==chars[i] && chars[j] != "" {
          cnt++
          chars[j]=""
        }
      }
      if (cnt > 1) {
        duplicates++
      }
    }
    
    return duplicates
    }