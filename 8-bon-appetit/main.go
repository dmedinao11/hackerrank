package main

import "strconv"

/*
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */

func bonAppetit(bill []int32, k int32, b int32) string {
	billToSplitSum := getSliceSum(append(bill[:k], bill[k+1:]...))
	billIndividualPart := billToSplitSum / 2

	if billIndividualPart == b {
		return "Bon Appetit"
	}

	return strconv.Itoa(int(b - billIndividualPart))
}

func getSliceSum(nums []int32) int32 {
	var result int32

	for _, num := range nums {
		result += num
	}

	return result
}
