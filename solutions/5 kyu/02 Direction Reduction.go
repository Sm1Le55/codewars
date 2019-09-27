package kata

func DirReduc(arr []string) []string {
	var oppDirect map[string]int = map[string]int{
		"NORTH": 1,	"SOUTH": -1, "EAST": 2, "WEST": -2,
	}

	for i := 0; i < len(arr)-1; {
		if oppDirect[arr[i]] + oppDirect[arr[i+1]] == 0 {
			arr = append(arr[:i], arr[i+2:]...)
			i = 0
		} else {
			i++
		}
	}
  if len(arr) == 0 {
    arr = append(arr,"")
  }
	return arr
}