package main

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (ac int32, oc int32) {
	ac = getFallenFruit(s, t, a, apples)
	oc = getFallenFruit(s, t, b, oranges)
	return
}

func getFallenFruit(s int32, t int32, treePoint int32, fruitsDistances []int32) (c int32) {
	for _, d := range fruitsDistances {
		position := treePoint + d
		if position >= s && position <= t {
			c++
		}
	}

	return
}
