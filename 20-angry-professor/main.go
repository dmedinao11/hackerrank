package main

/*
 * Complete the 'angryProfessor' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY a
 */

const (
	yes = "YES"
	no  = "NO"
)

func angryProfessor(k int32, a []int32) string {

	var count int32

	for _, i := range a {
		if i <= 0 {
			count++
		}

		if count >= k {
			return no
		}
	}

	return yes
}
