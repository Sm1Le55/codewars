package kata

func FindUniq(arr []float32) float32 {
	var repeatValue float32
	if arr[0] == arr[1] {
		repeatValue = arr[0]
	} else if arr[0] == arr[2] {
		return arr[1]
	} else {
		return arr[0]
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] != repeatValue {
			return arr[i]
		}
	}
	return repeatValue
}
