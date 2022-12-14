package main

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	var lcm int32

	for _, i := range a {
		if lcm == 0 {
			lcm = i
			continue
		}

		lcm = GetLCM(lcm, i)
	}

	var gcd int32

	for _, j := range b {
		if gcd == 0 {
			gcd = j
			continue
		}

		gcd = GetGCD(gcd, j)
	}

	var count int32

	for i := lcm; i <= gcd; i += lcm {
		if gcd%i == 0 {
			count++
		}
	}

	return count
}

func GetGCD(a, b int32) int32 {
	if b == 0 {
		return a
	}
	return GetGCD(b, a%b)
}

func GetLCM(a, b int32) int32 {
	result := a * b / GetGCD(a, b)

	return result
}
