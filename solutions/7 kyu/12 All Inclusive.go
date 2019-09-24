package kata

func ContainAllRots(strng string, arr []string) bool { 
    for _,c := range strng {
    strng = strng[1:]+string(c)
    isFind := false
      for _, val := range arr {
        if val==strng {
          isFind = true
          break        
        } 
      }
    if !isFind {
      return false
      }
    }
    return true
}