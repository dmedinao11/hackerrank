package main

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

var allMagicSquares = [][][]int32{
	{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
}

func formingMagicSquare(s [][]int32) int32 {
	var result int32

	for i, square := range allMagicSquares {
		squareCost := calculateSquareConversionCost(s, square)

		if i == 0 || squareCost < result {
			result = squareCost
		}
	}

	return result
}

func calculateSquareConversionCost(a [][]int32, b [][]int32) int32 {
	var result int32

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result += abs(a[i][j] - b[i][j])
		}
	}

	return result
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
