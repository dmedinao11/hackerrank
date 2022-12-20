package main

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	var countMapIndex [100]int32
	var maxSubsetLength int32

	for _, i := range a {
		countMapIndex[i]++
	}

	for i := 2; i < 100; i++ {
		consecutiveSum := countMapIndex[i] + countMapIndex[i-1]
		if consecutiveSum > maxSubsetLength {
			maxSubsetLength = consecutiveSum
		}
	}

	return maxSubsetLength
}
