package binary_tree

type (
	Node[C any] struct {
		Value C
		Left  *Node[C]
		Right *Node[C]
	}
)

func NewNode[C any](value C) Node[C] {
	return Node[C]{Value: value}
}
