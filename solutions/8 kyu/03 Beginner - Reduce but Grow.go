package kata

func Grow(arr []int) int{
result := 1
for _, val := range arr{
  result = result * val
}
return result
}