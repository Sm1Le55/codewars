package kata

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) (result string) {
	numbers := strings.Split(in, " ")
	min, _ := strconv.Atoi(numbers[0])
	max, _ := strconv.Atoi(numbers[0])

	for _, val := range numbers {
		num, _ := strconv.Atoi(val)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	result = strconv.Itoa(max) + " " + strconv.Itoa(min)
	return
}
