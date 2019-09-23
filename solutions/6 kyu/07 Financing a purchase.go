package kata

import (
	"fmt"
	"math"
)

func Amort(rate float64, p int, n int, num_payment int) string {

	r := rate / 12 / 100
	c := (r * float64(p)) / (1 - math.Pow(1+r, float64(-n)))

	x := math.Pow((1 + r), float64(num_payment))
	xPrev := math.Pow((1 + r), float64(num_payment-1))

	balance := x*float64(p) - (((x - 1) / r) * c)
	balancePrev := xPrev*float64(p) - (((xPrev - 1) / r) * c)

	interest := r * float64(balancePrev)
	princ := c - interest

	result := fmt.Sprintf("num_payment %d c %.0f princ %.0f int %.0f balance %.0f", num_payment, c, princ, interest, balance)
	fmt.Println(result)
	return result
}
