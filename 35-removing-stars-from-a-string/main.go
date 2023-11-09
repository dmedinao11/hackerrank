package main

func removeStars(s string) string {
	result := ""
	starCount := 0

	for i := len(s) - 1; i >= 0; i-- {
		char := string(s[i])
		if char == "*" {
			starCount++
			continue
		}

		if starCount == 0 {
			result = char + result
			continue
		}

		starCount--
	}

	return result
}
