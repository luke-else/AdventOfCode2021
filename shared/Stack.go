package shared

type Stack struct {
	head *Node
}

func (s *Stack) Push(item string) {
	if s.head == nil {
		s.head = &Node{
			Value: item,
			Next:  nil,
		}
	} else {
		s.head = &Node{
			Value: item,
			Next:  s.head,
		}
	}
}

func (s *Stack) Pop() string {
	node := s.Peek()
	s.head = node.Next
	return node.Value
}

func (s *Stack) Peek() (node *Node) {
	node = s.head
	if node == nil {
		node = &Node{}
	}
	return
}
