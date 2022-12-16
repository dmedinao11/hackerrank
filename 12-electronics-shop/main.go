package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maxB int32 = -1

	for _, l := range keyboards {
		if l >= b {
			continue
		}

		for _, s := range drives {
			sum := l + s

			if sum == b {
				return b
			}

			if sum > maxB && sum <= b {
				maxB = sum
			}
		}
	}

	return maxB
}

func convertToInt(ar []int32) []int {
	newArray := make([]int, len(ar))
	for i, v := range ar {
		newArray[i] = int(v)
	}
	return newArray
}

func main() {
	a := convertToInt([]int32{1, 3, 2})
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}
