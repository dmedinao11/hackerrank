package main

import "fmt"

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	const (
		twelve    = 12
		thirteen  = 13
		twentySix = 26
	)

	result := thirteen

	if year%4 == 0 && (year < 1918 || year%400 == 0 || year%100 != 0) {
		result = twelve
	}

	if year == 1918 {
		result = twentySix
	}

	return fmt.Sprintf("%d.09.%d", result, year)
}
