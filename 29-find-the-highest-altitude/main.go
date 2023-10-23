package main

func largestAltitude(gain []int) int {
	var max, sum int

	for i := 0; i < len(gain); i++ {
		if ca := gain[i] + sum; ca > max {
			max = ca
		}

		sum += gain[i]
	}

	return max
}
