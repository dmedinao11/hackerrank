package main

type (
	node struct {
		value int
		next  *node
	}

	stack struct {
		head *node
	}
)

func (s *stack) push(value int) {
	n := &node{value: value, next: s.head}
	s.head = n
}

func (s *stack) pop() *int {
	if s.head == nil {
		return nil
	}

	head := s.head
	s.head = head.next
	return &head.value
}

func (s *stack) peek() *int {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}

func (s *stack) toSlice() []int {
	result := make([]int, 0)
	currentNode := s.head

	for currentNode != nil {
		result = append(result, currentNode.value)
		currentNode = currentNode.next
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func asteroidCollision(asteroids []int) []int {
	s := stack{}

	for _, asteroid := range asteroids {
		if asteroid > 0 || s.peek() == nil || *s.peek() < 0 {
			s.push(asteroid)
			continue
		}

		abs := asteroid * -1
		for s.peek() != nil {
			topValue := *s.peek()

			if topValue < 0 {
				s.push(asteroid)
				break
			}

			if topValue > abs {
				break
			}

			s.pop()

			if topValue == abs {
				break
			}

			if s.peek() == nil {
				s.push(asteroid)
				break
			}
		}
	}

	return s.toSlice()
}
