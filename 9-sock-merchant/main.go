package main

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	var pairsMap [101]bool
	var count int32

	for _, i := range ar {
		if pairsMap[i] {
			pairsMap[i] = false
			count++
		} else {
			pairsMap[i] = true
		}
	}

	return count
}
