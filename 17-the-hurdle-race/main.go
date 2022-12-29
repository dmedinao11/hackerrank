package main

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

func hurdleRace(k int32, height []int32) int32 {
	var maxValue = getMaxValue(height)

	if k > maxValue {
		return 0
	} else {
		return maxValue - k
	}
}

func getMaxValue(height []int32) int32 {
	if height == nil || len(height) == 0 {
		return 0
	}

	max := height[0]

	for _, i := range height {
		if i > max {
			max = i
		}
	}

	return max
}
