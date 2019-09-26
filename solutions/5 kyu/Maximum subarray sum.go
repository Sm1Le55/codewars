package kata

func MaximumSubarraySum(numbers []int) int {
	maxSum := 0

	for cnt := 1; cnt < len(numbers); cnt++ {
		for i := 0; i < len(numbers) - cnt + 1; i++ {
			sum := 0
			for j := i; j < i + cnt; j++ {
				sum += numbers[j]
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}