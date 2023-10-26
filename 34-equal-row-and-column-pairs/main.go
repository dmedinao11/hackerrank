package main

import "strconv"

func equalPairs(grid [][]int) int {
	columnsMap := make(map[string]int, len(grid))
	response := 0

	for i := 0; i < len(grid); i++ {
		result := ""
		for j := 0; j < len(grid); j++ {
			result += strconv.Itoa(grid[j][i]) + "_"
		}
		columnsMap[result]++
	}

	for _, nums := range grid {
		result := ""
		for _, num := range nums {
			result += strconv.Itoa(num) + "_"
		}

		if columnsMap[result] != 0 {
			response += columnsMap[result]
		}
	}

	return response
}
