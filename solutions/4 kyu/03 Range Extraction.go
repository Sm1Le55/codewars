package kata

import (
	"math"
	"strconv"
)

//Solution ...
func Solution(list []int) (result string) {
	line := 0

	for index, element := range list {

		//Processing first element
		if index == 0 {
			result += strconv.Itoa(element)
			continue
		}

		//Processing equal elements
		if list[index-1] == element {
			continue
		}

		//Check that the elements are nearby in the number series.
		if element-list[index-1] == int(math.Abs(1)) {
			line++
			//Processing last element
			if index == len(list)-1 {
				if line > 1 {
					result += "-" + strconv.Itoa(element)
				} else {
					result += "," + strconv.Itoa(element)
				}
			}
		} else {
			if line > 1 {
				//Processing last element
				if index == len(list)-1 {
					result += "-" + strconv.Itoa(list[index-1]) + "," + strconv.Itoa(element)
				} else {
					result += "-" + strconv.Itoa(list[index-1]) + "," + strconv.Itoa(element)
				}
			} else if line == 1 {
				result += "," + strconv.Itoa(list[index-1]) + "," + strconv.Itoa(element)
			} else {
				result += "," + strconv.Itoa(element)
			}
			line = 0
		}

	}

	return
}
