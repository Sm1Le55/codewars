package kata
 
import "strings"
	
func ToNato(words string) (result string) {
	buf := []string{}

	for _, ch := range words {
		if ch == ' ' {
			continue
		}
		buf = append(buf, searchNatoWord(string(ch)))
	}

	result = strings.Join(buf," ")
	return
}

func searchNatoWord(ch string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot","Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}	
	for _, n := range nato {
		if strings.ToUpper(string(ch)) == n[:1] {
			return n
		}
	}
	return ch
}