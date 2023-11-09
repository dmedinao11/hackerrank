package main

import (
	binaryTree "github.com/dmedinao11/hackerrank/binary-tree"
	"reflect"
	"testing"
)

func Test_depthFirstValues(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		// Given
		a := binaryTree.NewNode("a")
		b := binaryTree.NewNode("b")
		c := binaryTree.NewNode("c")
		d := binaryTree.NewNode("d")
		e := binaryTree.NewNode("e")
		f := binaryTree.NewNode("f")

		a.Left = &b
		a.Right = &c
		b.Left = &d
		b.Right = &e
		c.Right = &f

		expected := []string{"a", "b", "d", "e", "c", "f"}

		// When
		got := depthFirstValues(a)

		// Then
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("asteroidCollision() = %v, want %v", got, expected)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		// Given
		a := binaryTree.NewNode("a")
		b := binaryTree.NewNode("b")
		c := binaryTree.NewNode("c")
		d := binaryTree.NewNode("d")
		e := binaryTree.NewNode("e")
		f := binaryTree.NewNode("f")
		g := binaryTree.NewNode("g")

		a.Left = &b
		a.Right = &c
		b.Left = &d
		b.Right = &e
		c.Right = &f
		e.Left = &g

		expected := []string{"a", "b", "d", "e", "g", "c", "f"}

		// When
		got := depthFirstValues(a)

		// Then
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("asteroidCollision() = %v, want %v", got, expected)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		// Given
		a := binaryTree.NewNode("a")
		b := binaryTree.NewNode("b")
		c := binaryTree.NewNode("c")
		d := binaryTree.NewNode("d")
		e := binaryTree.NewNode("e")

		a.Right = &b
		b.Left = &c
		c.Right = &d
		d.Right = &e

		expected := []string{"a", "b", "c", "d", "e"}

		// When
		got := depthFirstValues(a)

		// Then
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("asteroidCollision() = %v, want %v", got, expected)
		}
	})
}
