package main

func prefixSumArray(arr []int) []int {
	var result []int

	result = append(result, arr[0])

	for i := 1; i < len(arr); i++ {
		result = append(result, result[i-1]+arr[i])
	}

	return result
}
