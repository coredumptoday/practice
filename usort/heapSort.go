package usort

import "github.com/coredumptoday/practice/heap"

func HeapSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	h := heap.BuildHeapFromSlice(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		v, _ := h.Pop()
		arr[i] = v
	}
}
