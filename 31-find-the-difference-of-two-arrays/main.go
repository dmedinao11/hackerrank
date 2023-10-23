package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool, len(nums1))
	set2 := make(map[int]bool, len(nums2))

	for _, num := range nums1 {
		set1[num] = true
	}

	result2 := make([]int, 0)

	for _, num := range nums2 {
		if !set1[num] {
			result2 = append(result2, num)
			set1[num] = true
		}
		set2[num] = true
	}

	result1 := make([]int, 0)

	for _, num := range nums1 {
		if !set2[num] {
			result1 = append(result1, num)
			set2[num] = true
		}
	}

	return [][]int{result1, result2}
}

//func findDifference(nums1 []int, nums2 []int) [][]int {
//	map2 := toMapSet(nums2)
//	res1 := calculateDifference(nums1, map2)
//
//	map1 := toMapSet(nums1)
//	res2 := calculateDifference(nums2, map1)
//
//	return [][]int{res1, res2}
//}
//
//func calculateDifference(nums []int, mapSet map[int]bool) []int {
//	addedMapSet := map[int]bool{}
//	result := make([]int, 0)
//
//	for _, num := range nums {
//		if !mapSet[num] && !addedMapSet[num] {
//			result = append(result, num)
//			addedMapSet[num] = true
//		}
//	}
//
//	return result
//}
//
//func toMapSet(nums1 []int) map[int]bool {
//	mapSet := map[int]bool{}
//	for _, num := range nums1 {
//		mapSet[num] = true
//	}
//	return mapSet
//}
