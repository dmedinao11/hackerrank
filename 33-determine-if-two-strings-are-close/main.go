package main

import (
	"reflect"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	charMap1 := map[rune]int{}

	for _, char := range word1 {
		charMap1[char]++
	}

	charMap2 := map[rune]int{}

	for _, char := range word2 {
		if charMap1[char] == 0 {
			return false
		}
		charMap2[char]++
	}

	occurrencesSorted1 := getOccurrences(charMap1)
	occurrencesSorted2 := getOccurrences(charMap2)

	return reflect.DeepEqual(occurrencesSorted1, occurrencesSorted2)
}

func getOccurrences(occurrences map[rune]int) []int {
	result := make([]int, len(occurrences))

	for _, count := range occurrences {
		result = append(result, count)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
