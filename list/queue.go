package list

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

func NewStack() *MyStack {
	s := &MyStack{
		queue: &doubleEndsQueue{
			head:   nil,
			tail:   nil,
			length: 0,
		},
	}
	return s
}

type MyStack struct {
	queue *doubleEndsQueue
}

func (s *MyStack) Push(value int32) {
	s.queue.addFromHead(value)
}

func (s *MyStack) Pop() *DoubleNode {
	return s.queue.popFromHead()
}

func (s *MyStack) IsEmpty() bool {
	return s.queue.isEmpty()
}

func (s *MyStack) Len() int32 {
	return s.queue.getLength()
}

func NewQueue() *MyQueue {
	s := &MyQueue{
		queue: &doubleEndsQueue{
			head:   nil,
			tail:   nil,
			length: 0,
		},
	}
	return s
}

type MyQueue struct {
	queue *doubleEndsQueue
}

func (s *MyQueue) Push(value int32) {
	s.queue.addFromHead(value)
}

func (s *MyQueue) Pop() *DoubleNode {
	return s.queue.popFromBottom()
}

func (s *MyQueue) IsEmpty() bool {
	return s.queue.isEmpty()
}

func (s *MyQueue) Len() int32 {
	return s.queue.getLength()
}
