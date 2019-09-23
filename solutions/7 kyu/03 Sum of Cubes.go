package kata

func SumCubes(n int) (result int) {
	for i := 0; i <= n; i++ {
		result += i * i * i
	}
	return
}
