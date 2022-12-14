package main

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	var counts [6]int32

	for _, a := range arr {
		counts[a]++
	}

	result := 1
	maxCount := counts[1]

	for i := 2; i < len(counts); i++ {
		if counts[i] > maxCount {
			result = i
			maxCount = counts[i]
		}
	}

	return int32(result)
}
