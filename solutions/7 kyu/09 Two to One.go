package kata

import (
  "strings"
  "sort"
)


func TwoToOne(s1 string, s2 string) string {
  slice1 := strings.Split(s1, "")
  slice2 := strings.Split(s2, "")
  slice1 = append(slice1, slice2...)
 
  sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j]})
    
  for i:=0; i<len(slice1)-1; i++ {  
    for j:=i+1; j<len(slice1); j++ {
      if slice1[i]==slice1[j] {
        slice1[j]=""
      }
    }
  }
  
  return strings.Join(slice1, "")
}