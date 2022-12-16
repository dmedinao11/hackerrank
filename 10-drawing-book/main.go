package main

import "fmt"

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
	turnCount := p / 2
	bookHalf := n / 2

	if p > bookHalf {
		turnCount = bookHalf - turnCount
	}

	return turnCount
}

func main() {
	a := 5
	b := 2
	c := a / b

	fmt.Printf("%d | %T \n", c, c)
}
