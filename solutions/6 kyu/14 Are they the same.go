package kata

func Comp(a []int, b []int) bool {

	if a == nil || b == nil {
		return false
	}

	if len(a) == 0 && len(b) == 0 {
		return true
	}

	for i := range b {
		isFind := false
		for j := range a {
			if a[j] != 0 && b[i] == (a[j]*a[j]) {
				a[j], b[i] = 0, 0
				isFind = true
				break
			}
		}
		if !isFind {
			return false
		}
	}
	return true
}
