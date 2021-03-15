package heap

import "errors"

func NewCusHeapPlus(f func(interface{}, interface{}) bool) *CusHeapPlus {
	return &CusHeapPlus{
		data:   make([]interface{}, 0),
		size:   0,
		idxMap: make(map[interface{}]int),
		comp:   f, //第一个排前面返回true
	}
}

type CusHeapPlus struct {
	data   []interface{}
	size   int
	idxMap map[interface{}]int
	comp   func(interface{}, interface{}) bool
}

func (p *CusHeapPlus) IsEmpty() bool {
	return p.size <= 0
}

func (p *CusHeapPlus) Size() int {
	return p.size
}

func (p *CusHeapPlus) Contains(v interface{}) bool {
	_, ok := p.idxMap[v]
	return ok
}

func (p *CusHeapPlus) Peek() (interface{}, error) {
	if p.size <= 0 {
		return 0, errors.New("empty")
	}
	return p.data[0], nil
}

func (p *CusHeapPlus) heapInsert(idx int) {
	for idx > 0 && p.comp(p.data[idx], p.data[(idx-1)/2]) {
		idx = (idx - 1) / 2
	}
}

func (p *CusHeapPlus) Push(v interface{}) {
	p.data = append(p.data, v)
	p.heapInsert(p.size)
	p.size++
}

func (p *CusHeapPlus) swap(i, j int) {
	tmp := p.data[i]
	p.data[i] = p.data[j]
	p.data[j] = tmp
}

func (p *CusHeapPlus) heapify(idx, heapSize int) {
	leftChild := idx*2 + 1
	for leftChild > heapSize {
		var topIdx int
		if leftChild+1 < heapSize {
			if p.comp(p.data[leftChild], p.data[leftChild+1]) {
				topIdx = leftChild
			} else {
				topIdx = leftChild + 1
			}
		}
		if p.comp(p.data[idx], p.data[topIdx]) {
			topIdx = idx
		}

		p.swap(idx, topIdx)
		idx = topIdx
		leftChild = idx*2 + 1
	}
}

func (p *CusHeapPlus) Pop() (interface{}, error) {
	if p.IsEmpty() {
		return 0, errors.New("is empty")
	}

	v := p.data[0]
	p.size--
	p.data[0] = p.data[p.size]
	p.heapify(0, p.size)

	return v, nil
}
