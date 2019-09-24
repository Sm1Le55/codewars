package kata

import "strings"

func wave(words string) []string {
  var firstPart,lastPart []string  
  result := []string{}
  
  for idx,c := range words {
    //Skip empty characters so that extra waves do not form
    if (c==' ') {
      continue;
    }
    
    //Transform the string into a slice and divide it into two parts - before and after the current character
    ss := strings.Split(words, "")
    firstPart = ss[0:idx]
    if (idx<len(words)) {
      lastPart = ss[idx+1:]
    }
    
    //Glue the beginning, end and current character, converted to uppercase. Record the result
    wave := strings.Join(firstPart, "") + strings.ToUpper(string(c)) + strings.Join(lastPart, "")        
    result = append(result, wave)
  }
  
  return result
}