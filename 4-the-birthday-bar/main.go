package main

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	var count int32

	for i := 0; i+int(m) <= len(s); i++ {
		if sum := getSliceSum(s[i : i+int(m)]); sum == d {
			count++
		}
	}

	return count
}

func getSliceSum(nums []int32) int32 {
	var result int32

	for _, num := range nums {
		result += num
	}

	return result
}
