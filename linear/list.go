package linear

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
	if head == nil {
		return nil
	}

	for head.Value == value {
		head = head.Next
		if head == nil {
			break
		}
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

type doubleEndsQueue struct {
	head   *DoubleNode
	tail   *DoubleNode
	length int32
}

func (q *doubleEndsQueue) isEmpty() bool {
	return q.head == nil
}

func (q *doubleEndsQueue) getLength() int32 {
	return q.length
}

func (q *doubleEndsQueue) addFromHead(value int32) {
	node := DoubleNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}

	if q.head == nil {
		q.tail = &node
	} else {
		node.Next = q.head
		q.head.Prev = &node
	}
	q.head = &node
	q.length++
}

func (q *doubleEndsQueue) addFromBottom(value int32) {
	node := DoubleNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}

	if q.head == nil {
		q.head = &node
	} else {
		node.Prev = q.tail
		q.tail.Next = &node
	}
	q.tail = &node
	q.length++
}

func (q *doubleEndsQueue) popFromHead() *DoubleNode {
	if q.head == nil {
		return nil
	}

	cur := q.head

	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}

	q.length--
	return cur
}

func (q *doubleEndsQueue) popFromBottom() *DoubleNode {
	if q.tail == nil {
		return nil
	}

	cur := q.tail

	q.tail = q.tail.Prev
	if q.tail == nil {
		q.head = nil
	} else {
		q.tail.Next = nil
	}

	q.length--
	return cur
}

//单链表返回成环点
func GetLoopNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if fast != slow {
		return nil
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
