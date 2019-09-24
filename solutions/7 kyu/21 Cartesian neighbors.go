package kata

func CartesianNeighbor(x,y int) [][]int{
  result := make([][]int, 0, 8)
  xAxis := []int{x-1,x,x+1}
  yAxis := []int{y-1,y,y+1}
  
  for _,valX := range xAxis {
    for _,valY := range yAxis {
      if (valX==x && valY==y) {
        continue
      }
      result = append(result,[]int{valX,valY})      
    }
  }
  
  return result
}