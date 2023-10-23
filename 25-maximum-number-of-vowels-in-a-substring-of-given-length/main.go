package main

var vowels = map[uint8]bool{
	97:  true,
	101: true,
	105: true,
	111: true,
	117: true,
}

func maxVowels(s string, k int) int {
	maxVowelsCount := 0

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			maxVowelsCount += 1
		}
	}

	currentVowels := maxVowelsCount

	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			currentVowels += 1
		}

		if vowels[s[i-k]] {
			currentVowels -= 1
		}

		if currentVowels > maxVowelsCount {
			maxVowelsCount = currentVowels
		}
	}

	return maxVowelsCount
}
