package kata

func Gps(s int, x []float64) (max int) {
  for i := 0; i < len(x) - 1; i++ {
    sectionSpeed := (3600 * (x[i+1] - x[i]) / float64(s))
    if sectionSpeed > float64(max) {
      max = int(sectionSpeed)
    }
  }
  return
}