package main

import "sort"

func uniqueOccurrences(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	occurrences := map[int]bool{}
	currentNum := arr[0]
	currentCount := 0

	for _, num := range arr {
		if num == currentNum {
			currentCount++
			continue
		}

		if occurrences[currentCount] {
			return false
		}

		occurrences[currentCount] = true
		currentCount = 1
		currentNum = num
	}

	return !occurrences[currentCount]
}
