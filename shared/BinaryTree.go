package shared

import (
	"fmt"
	"strconv"
)

type BinaryTree struct {
	Head *Node
}

func (t *BinaryTree) Insert(val int) {
	if t.Head == nil {
		t.Head = &Node{Value: strconv.Itoa(val), Left: nil, Right: nil}
		return
	}

	current := t.Head
	var parent *Node = nil

	for {
		currentVal, _ := strconv.Atoi(current.Value)
		parent = current
		left := true
		if val > currentVal {
			left = false
			_, current = t.Traverse(current)
		} else {
			current, _ = t.Traverse(current)
		}

		if current == nil {
			current = &Node{Value: strconv.Itoa(val), Left: nil, Right: nil, Parent: parent}
			if left {
				parent.Left = current
			} else {
				parent.Right = current
			}
			return
		}
	}
}

func (t *BinaryTree) InOrder(start *Node) {
	if start.Left != nil {
		t.InOrder(start.Left)
	}
	fmt.Print(start.Value, " ")
	if start.Right != nil {
		t.InOrder(start.Right)
	}
}

func (t *BinaryTree) PreOrder(start *Node) {
	fmt.Print(start.Value, " ")
	if start.Left != nil {
		t.PreOrder(start.Left)
	}
	if start.Right != nil {
		t.PreOrder(start.Right)
	}
}

func (t *BinaryTree) PostOrder(start *Node) {
	if start.Left != nil {
		t.PostOrder(start.Left)
	}
	if start.Right != nil {
		t.PostOrder(start.Right)
	}
	fmt.Print(start.Value, " ")
}

func (t *BinaryTree) Traverse(current *Node) (left *Node, right *Node) {
	return current.Left, current.Right
}
