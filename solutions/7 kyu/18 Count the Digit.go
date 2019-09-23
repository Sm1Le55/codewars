package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	fmt.Println("n", n, "d", d)
	allSquares := ""
	for i := 0; i <= n; i++ {
		allSquares += strconv.Itoa(i * i)
	}
	return strings.Count(allSquares, strconv.Itoa(d))
}
