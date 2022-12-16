package main

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	const (
		u = 85
	)

	var (
		lastHigh int
		high     int
		count    int32
	)

	for _, s := range path {
		lastHigh = high

		if s == u {
			high++
		} else {
			high--
		}

		if lastHigh == 0 && high == -1 {
			count++
		}
	}

	return count
}
