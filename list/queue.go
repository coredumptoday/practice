package list

type DoubleEndsQueue struct {
	head   *DoubleNode
	tail   *DoubleNode
	length int32
}

func (q *DoubleEndsQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *DoubleEndsQueue) GetLength() int32 {
	return q.length
}

func (q *DoubleEndsQueue) AddFromHead(value int32) {
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

func (q *DoubleEndsQueue) AddFromBottom(value int32) {
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

func (q *DoubleEndsQueue) PopFromHead() *DoubleNode {
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

func (q *DoubleEndsQueue) PopFromBottom() *DoubleNode {
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
