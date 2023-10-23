package main

func longestOnes(nums []int, k int) int {
	var max, i, j, zc int

	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			zc++
		}

		if zc > k {
			count := i - j
			if count > max {
				max = count
			}

			if nums[j] == 0 {
				zc--
			}

			j++
		}
	}

	count := i - j
	if count > max {
		max = count
	}

	return max
}
