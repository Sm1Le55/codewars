package kata

func IsTriangle(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	maxSide := a
	if b > maxSide {
		maxSide = b
	}
	if c > maxSide {
		maxSide = c
	}

	if a+b+c-maxSide > maxSide {
		return true
	} else {
		return false
	}

}
