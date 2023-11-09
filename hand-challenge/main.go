package main

import (
	"strings"
)

func translate(value string) string {
	index := 0
	memory := map[int]int{}
	memoryPointer := 0
	instructions := strings.Split(value, "")
	result := ""

	for index < len(instructions) {
		instruction := instructions[index]

		switch instruction {
		case "👉":
			memoryPointer++
		case "👈":
			memoryPointer--
		case "👆":
			memory[memoryPointer] = (memory[memoryPointer] + 1) % 256
		case "👇":
			memoryValue := memory[memoryPointer] - 1

			if memoryValue == -1 {
				memory[memoryPointer] = 255
			} else {
				memory[memoryPointer] = memoryValue % 256
			}
		case "🤜":
			if memory[memoryPointer] == 0 {
				index = getCloseFistIndex(instructions, index)
			}
		case "🤛":
			if memory[memoryPointer] != 0 {
				index = getOpenFistIndex(instructions, index)
			}
		case "👊":
			result += string(rune(memory[memoryPointer]))
		}
		index++
	}
	return result
}

func getCloseFistIndex(instructions []string, index int) int {
	i := index + 1
	openFistCount := 0

	for ; i < len(instructions); i++ {
		if instructions[i] == "🤜" {
			openFistCount++
		}

		if instructions[i] == "🤛" {
			if openFistCount == 0 {
				return i
			}
			openFistCount--
		}
	}
	return index
}

func getOpenFistIndex(instructions []string, index int) int {
	i := index - 1
	closeFistCount := 0

	for ; i >= 0; i-- {
		if instructions[i] == "🤜" {
			if closeFistCount == 0 {
				return i
			}
			closeFistCount--
		}

		if instructions[i] == "🤛" {
			closeFistCount++
		}
	}
	return index
}
