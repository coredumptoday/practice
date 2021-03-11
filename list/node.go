package list

type Node struct {
	Value int32
	Next  *Node
}

type DoubleNode struct {
	Value int32
	Next  *DoubleNode
	Prev  *DoubleNode
}

func RevLinkedList(head *Node) *Node {
	var prev *Node
	var next *Node

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
