package kata

import (
  "strings"
  "sort"
  "strconv"
)

func Hist(s string) (result string) {
    //Объявили список ошибок
    errorsList := [...]rune{'u','w','x','z'}
    //Соритируем внутренности строки
    str := strings.Split(s, "")
    sort.Slice(str, func(i, j int) bool { return str[i] < str[j]})   
    s = strings.Join(str, "")
    
    //Открываем цикл по списку ошибок
    for _, printerError := range errorsList {
      cnt := 0      
      //Обходим в цикле строку
      for _, char := range s {       
        if char==printerError {        
          cnt++
        }
      }
      
      if cnt==0 {
        continue
      }
      
      if result!="" {
        result += "\r"
      }      
      result += string(printerError)
      result += "  "
      result += strconv.Itoa(cnt)
      result += "     "
      for i:=0; i<cnt; i++ {
        result += "*"
      }
    }
    return
}