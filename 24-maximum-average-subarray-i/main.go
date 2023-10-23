package main

func findMaxAverage(nums []int, k int) float64 {
	var maxSum int

	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	currentSum := maxSum

	for i := k; i < len(nums); i++ {
		currentSum = currentSum - nums[i-k] + nums[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return float64(maxSum) / float64(k)
}
