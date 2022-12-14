package main

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	if len(scores) <= 1 {
		return []int32{0, 0}
	}

	max := scores[0]
	min := scores[0]
	var countMax int32
	var countMin int32

	for _, score := range scores {
		if score > max {
			countMax++
			max = score
			continue
		}

		if score < min {
			countMin++
			min = score
		}
	}

	return []int32{countMax, countMin}
}
