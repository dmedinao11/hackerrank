package main

import "fmt"

func main() {
	fmt.Println(!canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1}, 1))
	fmt.Println(!canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); {
		if canBePlanted(flowerbed, i) {
			n -= 1

			if n == 0 {
				return true
			}

			i += 2
			continue
		}

		i += 1
	}

	return n == 0
}

func canBePlanted(flowerbed []int, i int) bool {
	if flowerbed[i] == 1 {
		return false
	}

	last := 0

	if i > 0 {
		last = flowerbed[i-1]
	}

	next := 0

	if i+1 <= len(flowerbed)-1 {
		next = flowerbed[i+1]
	}

	return next == 0 && last == 0
}
