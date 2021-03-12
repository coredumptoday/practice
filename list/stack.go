package list

import (
	clist "container/list"
)

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

func NewMinStack() *MinStack {
	return &MinStack{
		stack: clist.New(),
		min:   clist.New(),
	}
}

type MinStack struct {
	stack *clist.List
	min   *clist.List
}

func (ms *MinStack) Push(v int) {
	ms.stack.PushFront(v)

	mf := ms.min.Front()
	if mf == nil {
		ms.min.PushFront(v)
	} else {
		res, _ := mf.Value.(int)
		if res >= v {
			ms.min.PushFront(v)
		} else {
			ms.min.PushFront(res)
		}
	}
}

func (ms *MinStack) Pop() int {
	f := ms.stack.Front()
	mf := ms.min.Front()

	ms.stack.Remove(f)
	ms.min.Remove(mf)

	res, _ := f.Value.(int)
	return res
}

func (ms *MinStack) GetMin() int {
	mf := ms.min.Front()
	res, _ := mf.Value.(int)
	return res
}

func NewQueueStack() *QueueStack {
	return &QueueStack{
		stackQueue: clist.New(),
		helpQueue:  clist.New(),
	}
}

type QueueStack struct {
	stackQueue *clist.List
	helpQueue  *clist.List
}

func (qs *QueueStack) Push(v int) {
	qs.stackQueue.PushFront(v)
}

func (qs *QueueStack) Poll() (int, bool) {
	if qs.stackQueue.Len() <= 0 {
		return 0, false
	}

	for qs.stackQueue.Len() > 1 {
		v := qs.stackQueue.Back()
		qs.stackQueue.Remove(v)
		qs.helpQueue.PushFront(v.Value)
	}

	v := qs.stackQueue.Back()
	qs.stackQueue.Remove(v)

	tmp := qs.stackQueue
	qs.stackQueue = qs.helpQueue
	qs.helpQueue = tmp

	res, ok := v.Value.(int)
	return res, ok
}

func (qs *QueueStack) Peek() (int, bool) {
	if qs.stackQueue.Len() <= 0 {
		return 0, false
	}

	for qs.stackQueue.Len() > 1 {
		v := qs.stackQueue.Back()
		qs.stackQueue.Remove(v)
		qs.helpQueue.PushFront(v.Value)
	}

	v := qs.stackQueue.Back()
	qs.stackQueue.Remove(v)
	qs.helpQueue.PushFront(v.Value)

	tmp := qs.stackQueue
	qs.stackQueue = qs.helpQueue
	qs.helpQueue = tmp

	res, ok := v.Value.(int)
	return res, ok
}

func (qs *QueueStack) IsEmpty() bool {
	return qs.stackQueue.Len() == 0
}

func (qs *QueueStack) Len() int {
	return qs.stackQueue.Len()
}
