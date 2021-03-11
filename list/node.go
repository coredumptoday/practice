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

func RevDoubleList(head *DoubleNode) *DoubleNode {
	var prev *DoubleNode
	var next *DoubleNode

	for head != nil {
		next = head.Next
		head.Next = prev
		head.Prev = next
		prev = head
		head = next
	}

	return prev
}

func RemoveValue(head *Node, value int32) *Node {
	for head.Value == value {
		head = head.Next
	}

	var cur *Node = head
	var prev *Node = head

	for cur != nil {
		if cur.Value == value {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return head
}
