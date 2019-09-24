package kata

import "math"

func WallPaper(l, w, h float64) string {
    nums := [...]string{"zero","one","two","three","four",
      "five","six","seven","eight","nine", "ten","eleven",
      "twelve","thirteen","fourteen","fifteen","sixteen",
      "seventeen","eighteen","nineteen","twenty"}
    
    if (l==0 || w==0 || h==0) {
      return nums[0]
    }

    perimeter := 2*(l+w)
    lineCount := perimeter/0.52
    lineInRolls := 10/h
    rolls := math.Round(lineCount/lineInRolls*1.15 + 0.5)
    
    return nums[int(rolls)]
}