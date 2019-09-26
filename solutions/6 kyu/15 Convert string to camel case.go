package kata

import "strings"

func ToCamelCase(s string) (result string) {
	sl := []string{}
	
	if strings.Index(s,"_") > 0 {
		sl = strings.Split(s,"_")
	} else {
		sl = strings.Split(s,"-")
	}

	result += sl[0]
	for i := 1; i < len(sl); i++ {
		word := sl[i]
		result += (strings.ToUpper(word[:1]) + word[1:])
	}

	return 
}

//TODO: strings.Title()