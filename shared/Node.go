package shared

//Waiting for generic types to release,
//this will allow for a Node with a different value type

type Node struct {
	Value string

	//Universal
	Parent *Node

	//For LLs, Stacks, Queues...
	Next *Node

	//For Binary Trees...
	Left  *Node
	Right *Node
}
