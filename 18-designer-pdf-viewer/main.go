package main

/*
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */
const asciiSlice int32 = 97

func designerPdfViewer(h []int32, word string) int32 {
	var maxHigh int32

	for _, l := range word {
		if h[l-asciiSlice] > maxHigh {
			maxHigh = h[l-asciiSlice]
		}
	}

	return maxHigh * int32(len(word))
}
