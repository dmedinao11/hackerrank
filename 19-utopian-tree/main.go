package main

/*
 * Complete the 'utopianTree' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func utopianTree(n int32) int32 {
	var result int32

	for i := 0; i <= int(n); i++ {
		if i%2 == 0 {
			result++
			continue
		}

		result *= 2
	}

	return result
}
