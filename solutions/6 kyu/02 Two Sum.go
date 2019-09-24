package kata

func TwoSum(numbers []int, target int) [2]int {    
    for i := range numbers {
      var firstNumber = numbers[i]      
      for j := range numbers {
        if (i != j && firstNumber+numbers[j]==target) {
          return [...]int{i, j}    
        }
      }
    }
    return [...]int{0, 0}     
}