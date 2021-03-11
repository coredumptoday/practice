package list

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

func NewRingArr(cap int) *RingArr {
	return &RingArr{
		head: 0,
		tail: 0,
		data: make([]int, cap),
		len:  0,
		cap:  cap,
	}
}

type RingArr struct {
	head int
	tail int
	data []int
	len  int
	cap  int
}

func (r *RingArr) nextIdx(v int) int {
	if v+1 < r.cap {
		return v + 1
	} else {
		return 0
	}
}

func (r *RingArr) Push(v int) bool {
	if r.len >= r.cap {
		return false
	}

	r.len++
	r.data[r.head] = v
	r.head = r.nextIdx(r.head)
	return true
}

func (r *RingArr) Pop() (int, bool) {
	if r.len <= 0 {
		return 0, false
	}

	r.len--
	d := r.data[r.tail]
	r.tail = r.nextIdx(r.tail)

	return d, true
}

func (r *RingArr) Len() int {
	return r.len
}
