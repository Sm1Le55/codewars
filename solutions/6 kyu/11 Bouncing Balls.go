package kata

func BouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}

	motherSeeBall := 0

	for h > window {
		motherSeeBall++
		h = h * bounce
		if h > window {
			motherSeeBall++
		}
	}
	return motherSeeBall
}
