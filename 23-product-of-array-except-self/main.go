package main

func productExceptSelf(nums []int) []int {
	rightProducts := make([]int, len(nums)-1)
	counter := 1

	for j := len(nums) - 1; j > 0; j-- {
		counter *= nums[j]
		rightProducts[j-1] = counter
	}

	leftProducts := make([]int, 0)
	result := make([]int, 0)

	counter = 1

	for i := 0; i < len(nums); i++ {
		counter *= nums[i]
		leftProducts = append(leftProducts, counter)

		if i == 0 {
			result = append(result, rightProducts[0])
			continue
		}

		if i == len(nums)-1 {
			result = append(result, leftProducts[i-1])
			continue
		}

		result = append(result, rightProducts[i]*leftProducts[i-1])
	}

	return result
}
