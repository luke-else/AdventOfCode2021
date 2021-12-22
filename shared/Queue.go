package shared

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Pop() (node *Node) {
	node = nil
	if q.Head != nil {
		node = q.Head
		q.Head = q.Head.Next
		if q.Head == nil {
			q.Tail = nil
		}
	}
	return
}

func (q *Queue) Add(node *Node) {

	if q.Tail != nil {
		q.Tail.Next = node
	} else {
		q.Head = node
	}

	q.Tail = node
}
