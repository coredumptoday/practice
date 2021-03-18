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

func RevLinkedListForRange(head *Node, m, n int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	h1S := head
	h1d := head

	for h1d != nil {
		count++
		if count+1 == m {
			break
		}
		h1d = h1d.Next
	}
	h2S := h1d.Next
	h2d := h2S
	h1d.Next = nil
	count++

	var prev *Node
	var next *Node

	for h2d != nil && count <= n {
		next = h2d.Next
		h2d.Next = prev
		prev = h2d
		h2d = next
		count++
	}

	h1d.Next = prev
	h2S.Next = next

	return h1S
}

func revGroup(cur *Node) (start *Node, end *Node) {
	end = cur

	var prev, next *Node

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev, end
}

func RevLinkedListForGroupK(head *Node, k int) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prevStart, prevEnd *Node
	var groupStart, groupEnd *Node
	backStart := head
	cur := head
	count := 0

	for cur != nil {
		count++
		backStart = cur.Next
		if count == 1 {
			groupStart = cur
		} else if count == k {
			groupEnd = cur
			groupEnd.Next = nil
			s, e := revGroup(groupStart)
			if prevStart == nil {
				prevStart = s
				prevEnd = e
			} else {
				prevEnd.Next = s
				prevEnd = e
			}
			prevEnd.Next = backStart
			groupStart = nil
			groupEnd = nil
			count = 0
		}
		cur = backStart
	}

	return prevStart
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

//获取两个非成环链表相交点
func GetLinkListCrossPointNoLoop(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	c1, c2 := 0, 0
	cur1 := head1
	cur2 := head2

	for cur1 != nil {
		c1++
		cur1 = cur1.Next
	}

	for cur2 != nil {
		c2++
		cur2 = cur2.Next
	}

	if cur1 != cur2 {
		return nil
	}

	for head1 != nil && head2 != nil {
		if c1 == c2 {
			if head1 == head2 {
				return head1
			}
			c1--
			c2--
			head1 = head1.Next
			head2 = head2.Next
		} else if c1 > c2 {
			c1--
			head1 = head1.Next
		} else {
			c2--
			head2 = head2.Next
		}
	}

	return nil
}

//获取两个成环链表相交点
func GetLinkListCrossPointWithLoop(head1, head2, endNode1, endNode2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	if endNode1 == endNode2 {
		c1, c2 := 0, 0
		cur1 := head1
		cur2 := head2

		for cur1 != endNode1 {
			c1++
			cur1 = cur1.Next
		}

		for cur2 != endNode1 {
			c2++
			cur2 = cur2.Next
		}

		for head1 != endNode1 && head2 != endNode1 {
			if c1 == c2 {
				if head1 == head2 {
					return head1
				}
				c1--
				c2--
				head1 = head1.Next
				head2 = head2.Next
			} else if c1 > c2 {
				c1--
				head1 = head1.Next
			} else {
				c2--
				head2 = head2.Next
			}
		}

		return nil
	} else {
		cur := endNode1.Next
		for cur != endNode1 {
			if cur == endNode2 {
				return endNode2
			}
		}
		return nil
	}
}

func LinkListRemoveRange(cur *Node, m, n int) {
	if cur == nil {
		return
	}
	count := 0
	var prev, back *Node
	for cur != nil {
		count++
		if count+1 == m {
			prev = cur
		}
		if count == n {
			back = cur.Next
			break
		}
		cur = cur.Next
	}

	prev.Next = back
}

func RemoveDupNodeFromSortList(cur *Node) {
	if cur == nil {
		return
	}

	curVal := cur.Value
	curStart := cur

	for cur != nil {
		cur = cur.Next
		if cur != nil && cur.Value != curVal {
			curStart.Next = cur
			curVal = cur.Value
			curStart = cur
		}
	}
}
