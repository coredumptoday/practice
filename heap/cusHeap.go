package heap

import (
	"errors"
	"fmt"
	"github.com/coredumptoday/practice/utils"
)

//手写大根堆
func NewCusMaxHeap(cap int) *CusMaxHeap {
	return &CusMaxHeap{
		data:  make([]int, cap),
		limit: cap,
		size:  0,
	}
}

type CusMaxHeap struct {
	data  []int
	limit int
	size  int
}

func (ch *CusMaxHeap) IsEmpty() bool {
	return ch.size <= 0
}

func (ch *CusMaxHeap) IsFull() bool {
	return ch.size == ch.limit
}

func (ch *CusMaxHeap) Push(v int) error {
	if ch.IsFull() {
		return errors.New("is full")
	}
	ch.data[ch.size] = v
	ch.heapInsert(ch.size)
	ch.size++

	return nil
}

func (ch *CusMaxHeap) heapInsert(idx int) {
	for idx > 0 && ch.data[idx] > ch.data[(idx-1)/2] {
		utils.IntSliceSwap(ch.data, idx, (idx-1)/2)
		idx = (idx - 1) / 2
	}
}

func (ch *CusMaxHeap) Pop() (int, error) {
	if ch.IsEmpty() {
		return 0, errors.New("is empty")
	}
	v := ch.data[0]

	ch.size--
	ch.data[0] = ch.data[ch.size]
	ch.heapify(0, ch.size)

	return v, nil
}

func (ch *CusMaxHeap) heapify(idx int, heapSize int) {
	leftChild := idx*2 + 1
	for leftChild < heapSize {
		var largest int
		if leftChild+1 >= heapSize {
			largest = leftChild
		} else if ch.data[leftChild] < ch.data[leftChild+1] {
			largest = leftChild + 1
		} else {
			largest = leftChild
		}

		if ch.data[largest] < ch.data[idx] {
			break
		}

		utils.IntSliceSwap(ch.data, largest, idx)
		idx = largest
		leftChild = idx*2 + 1
	}
}

func (ch *CusMaxHeap) PrintSlice() {
	fmt.Println("MaxHeap: ", ch.data, "len: ", ch.size)
}

func BuildHeapFromSlice(arr []int) *CusMaxHeap {
	size := len(arr)

	h := &CusMaxHeap{
		data:  make([]int, size),
		limit: size,
		size:  size,
	}

	for i := size - 1; i >= 0; i-- {
		h.data[i] = arr[i]
		h.heapify(i, size)
	}

	heapSize := size
	heapSize--
	utils.IntSliceSwap(arr, 0, heapSize)
	// O(N*logN)
	for heapSize > 0 { // O(N)
		h.heapify(0, heapSize) // O(logN)
		heapSize--
		utils.IntSliceSwap(arr, 0, heapSize) // O(1)
	}

	return h
}

func NewHeapSlice(cap int) *MaxHeapSlice {
	return &MaxHeapSlice{
		data: make([]int, cap),
		cap:  cap,
		size: 0,
	}
}

type MaxHeapSlice struct {
	data []int
	cap  int
	size int
}

func (hs *MaxHeapSlice) IsEmpty() bool {
	return hs.size == 0
}

func (hs *MaxHeapSlice) IsFull() bool {
	return hs.size == hs.cap
}

func (hs *MaxHeapSlice) Push(v int) error {
	if hs.IsFull() {
		return errors.New("slice is full")
	}

	hs.data[hs.size] = v
	hs.size++

	return nil
}

func (hs *MaxHeapSlice) Pop() (int, error) {
	if hs.IsEmpty() {
		return 0, errors.New("slice is empty")
	}

	maxIdx := 0
	for i := 0; i < hs.size; i++ {
		if hs.data[i] > hs.data[maxIdx] {
			maxIdx = i
		}
	}

	v := hs.data[maxIdx]
	hs.size--
	hs.data[maxIdx] = hs.data[hs.size]

	return v, nil
}

func (hs *MaxHeapSlice) PrintSlice() {
	fmt.Println("MaxSlice: ", hs.data)
}
