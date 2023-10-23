package main

func longestSubarray(nums []int) int {
	var max, i, j int

	lz := -1

	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			if lz == -1 {
				lz = i
				continue
			}

			count := i - j
			if count > max {
				max = count
			}

			j = lz + 1
			lz = i
		}
	}

	count := i - j
	if count > max {
		max = count
	}

	return max - 1
}
