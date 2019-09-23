package kata

import "math"

func MxDifLg(a1 []string, a2 []string) (minLength int) {
    if len(a1) == 0 || len(a2) == 0 {
      return -1
    }
    for _, val1 := range a1 {
      for _, val2 := range a2 {
        if int(math.Abs(float64(len(val2) - len(val1)))) > minLength {
          minLength = int(math.Abs(float64(len(val2) - len(val1))))
        }
      }
    }
    return
}