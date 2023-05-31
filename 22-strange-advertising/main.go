package main

func viralAdvertising(n int32) int32 {
	var shared, result int32 = 5, 0

	for i := 0; i < int(n); i++ {
		result += shared / 2
		shared = shared / 2 * 3
	}

	return result
}
