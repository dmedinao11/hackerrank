package main

func pivotIndex(nums []int) int {
	var l, r int
	psa := prefixSumArray(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 {
			l = psa[i-1]
		}

		r = psa[len(psa)-1] - psa[i]

		if l == r {
			return i
		}
	}

	return -1
}

func prefixSumArray(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var result []int

	result = append(result, arr[0])

	for i := 1; i < len(arr); i++ {
		result = append(result, result[i-1]+arr[i])
	}

	return result
}
