package kata

const WIN_POINTS = 3
const DRAW_POINTS = 1

func Points(games []string) int {
  var totalPoints int
  for _, val := range games {
    x := string(val[:1])
    y := string(val[2:])
    
    if (x > y) {
      totalPoints += WIN_POINTS
    } else if (x == y) {
      totalPoints += DRAW_POINTS
    } 
  }
  return totalPoints
}