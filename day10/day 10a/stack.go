package main

type Stack struct {
	head *StackNode
}

func (s *Stack) Push(item byte) {
	if s.head == nil {
		s.head = &StackNode{
			char: item,
			next: nil,
		}
	} else {
		new := StackNode{
			char: item,
			next: s.head,
		}
		s.head = &new
	}
}

func (s *Stack) Pop() byte {
	node := s.Peek()
	s.head = node.next
	return node.char
}

func (s *Stack) Peek() (node *StackNode) {
	node = s.head
	if node == nil {
		node = &StackNode{}
	}
	return
}

type StackNode struct {
	char byte
	next *StackNode
}
