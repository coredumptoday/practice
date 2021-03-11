package list

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
