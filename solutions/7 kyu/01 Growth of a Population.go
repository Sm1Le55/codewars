package kata

import "math"

func NbYear(p0 int, percent float64, aug int, p int) (years int) {
	for p0 < p {
		p0 = int(math.Ceil(float64(p0) + float64(p0)*(percent/100) + float64(aug)))
		years++
	}
	return
}
