package stack

type (
	Stack[C any] interface {
		Push(value C)
		Pop() *C
		Peek() *C
		ToSlice() []C
	}

	stack[C any] struct {
		size int
		head *node[C]
	}

	node[C any] struct {
		value C
		next  *node[C]
	}
)

func NewStack[C any]() Stack[C] {
	return &stack[C]{}
}

func NewStackFromItems[C any](items ...C) Stack[C] {
	s := NewStack[C]()

	for _, item := range items {
		s.Push(item)
	}

	return s
}

func (s *stack[C]) Push(value C) {
	s.size++
	n := &node[C]{value: value, next: s.head}
	s.head = n
}

func (s *stack[C]) Pop() *C {
	if s.head == nil {
		return nil
	}

	head := s.head
	s.head = head.next
	s.size--
	return &head.value
}

func (s *stack[C]) Peek() *C {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}

func (s *stack[C]) ToSlice() []C {
	result := make([]C, 0, s.size)
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
