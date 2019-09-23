package kata

func Multiple3And5(number int) (total int) {
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return
}
