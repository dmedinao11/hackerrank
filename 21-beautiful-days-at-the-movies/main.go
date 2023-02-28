package main

import "strconv"

func beautifulDays(i int32, j int32, k int32) int32 {
	var result int32

	for n := i; n <= j; n++ {
		absoluteSub := getAbsoluteValue(n - reverseNumber(n))

		if absoluteSub%k == 0 {
			result += 1
		}
	}

	return result
}

func reverseNumber(n int32) int32 {
	strNumber := strconv.FormatInt(int64(n), 10)
	var strResult string

	for i := len(strNumber) - 1; i >= 0; i-- {
		if strResult == "" && string(strNumber[i]) == "0" {
			continue
		}

		strResult += string(strNumber[i])
	}

	result, _ := strconv.Atoi(strResult)

	return int32(result)
}

func getAbsoluteValue(n int32) int32 {
	if n < 0 {
		return n * -1
	}

	return n
}
