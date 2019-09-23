package kata

func InAscOrder(numbers []int) bool {
	if len(numbers) <= 1 {
		return true
	}
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			return false
		}
	}
	return true
}
