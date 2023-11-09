package main

import "fmt"

type CustomStack struct {
	maxSize int
	index   int
	array   []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		maxSize: maxSize,
		index:   0,
		array:   make([]int, maxSize),
	}
}

func (s *CustomStack) Push(x int) {
	if s.maxSize == s.index {
		return
	}
	s.array[s.index] = x
	s.index++
}

func (s *CustomStack) Pop() int {
	if s.index == 0 {
		return -1
	}

	s.index--
	return s.array[s.index]
}

func (s *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < s.index; i++ {
		s.array[i] += val
	}
}

func main() {
	stk := Constructor(3)
	stk.Push(1)
	stk.Push(2)
	fmt.Println(stk.Pop() == 2)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.Increment(5, 100)
	stk.Increment(2, 100)
	fmt.Println(stk.Pop() == 103)
	fmt.Println(stk.Pop() == 202)
	fmt.Println(stk.Pop() == 201)
	fmt.Println(stk.Pop() == -1)
}
