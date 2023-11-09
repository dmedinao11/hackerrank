package main

import (
	binaryTree "github.com/dmedinao11/hackerrank/binary-tree"
	"github.com/dmedinao11/hackerrank/stack"
)

func depthFirstValues[C any](root binaryTree.Node[C]) []C {
	itemsStack := stack.NewStackFromItems(root)
	result := make([]C, 0)

	for itemsStack.Peek() != nil {
		node := itemsStack.Pop()
		result = append(result, node.Value)

		if node.Right != nil {
			itemsStack.Push(*node.Right)
		}

		if node.Left != nil {
			itemsStack.Push(*node.Left)
		}
	}

	return result
}
