package shared

import "fmt"

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) InsertItem(first *Node, second *Node, value string) {
	first.Next = &Node{Value: value, Next: second}
}

func (l *LinkedList) PrintList() {
	list := []string{}
	current := l.Head
	for current != nil {
		list = append(list, current.Value)
		current = current.Next
	}
	fmt.Println(list)
}
